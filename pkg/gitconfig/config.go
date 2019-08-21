package gitconfig

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/pkg/errors"
)

const (
	gitconfigPath     = ".git/config"
	regexpGitHTTPSURL = `(?:https+):(\/\/)?(.*?)(\.git)(\/?|\#[-\d\w._]+?)$`
	regexpGitSSHURL   = `(?:git@[-\w.]+):(\/\/)?(.*?)(\.git)(\/?|\#[-\d\w._]+?)$`

	utf8BOM = "\357\273\277"
)

var (
	// ErrInvalidEscapeSequence indicates that the escape character ('\')
	// was followed by an invalid character.
	ErrInvalidEscapeSequence = errors.New("unknown escape sequence")

	// ErrUnfinishedQuote indicates that a value has an odd number of (unescaped) quotes
	ErrUnfinishedQuote = errors.New("unfinished quote")

	// ErrMissingEquals indicates that an equals sign ('=') was expected but not found
	ErrMissingEquals = errors.New("expected '='")

	// ErrPartialBOM indicates that the file begins with a partial UTF8-BOM
	ErrPartialBOM = errors.New("partial UTF8-BOM")

	// ErrInvalidKeyChar indicates that there was an invalid key character
	ErrInvalidKeyChar = errors.New("invalid key character")

	// ErrInvalidSectionChar indicates that there was an invalid character in section
	ErrInvalidSectionChar = errors.New("invalid character in section")

	// ErrUnexpectedEOF indicates that there was an unexpected EOF
	ErrUnexpectedEOF = errors.New("unexpected EOF")

	// ErrSectionNewLine indicates that there was a newline in section
	ErrSectionNewLine = errors.New("newline in section")

	// ErrMissingStartQuote indicates that there was a missing start quote
	ErrMissingStartQuote = errors.New("missing start quote")

	// ErrMissingClosingBracket indicates that there was a missing closing bracket in section
	ErrMissingClosingBracket = errors.New("missing closing section bracket")
)

type config struct {
	Branch       string
	Remote       string
	RemoteConfig *RemoteConfig
}

type RemoteConfig struct {
	URL          string
	Organization string
	Repository   string
}

type parser struct {
	bytes  []byte
	linenr uint
	eof    bool
}

func Config() (*config, error) {
	f, err := os.Open(gitconfigPath)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	m, _, err := Parse(b)
	if err != nil {
		return nil, err
	}

	var branch, remote string
	remoteConfig := &RemoteConfig{}
	for k, v := range m {
		splited := strings.Split(k, ".")
		prefix := splited[0]
		if prefix == "remote" && splited[2] == "url" {
			remoteConfig, err = NewRemoteConfig(v)
			if err != nil {
				return nil, err
			}
			continue
		}

		if prefix == "branch" && splited[2] == "remote" {
			branch = splited[1]
			remote = v
			continue
		}
	}

	return &config{
		Branch:       branch,
		Remote:       remote,
		RemoteConfig: remoteConfig,
	}, err
}

func Parse(bytes []byte) (map[string]string, uint, error) {
	parser := &parser{bytes, 1, false}
	cfg, err := parser.parse()
	return cfg, parser.linenr, err
}

func (cf *parser) parse() (map[string]string, error) {
	bomPtr := 0
	comment := false
	cfg := map[string]string{}
	name := ""
	var err error
	for {
		c := cf.nextChar()
		if bomPtr != -1 && bomPtr < len(utf8BOM) {
			if c == (utf8BOM[bomPtr] & 0377) {
				bomPtr++
				continue
			}
		} else {
			/* Do not tolerate partial BOM. */
			if bomPtr != 0 {
				return cfg, ErrPartialBOM
			}
			bomPtr = -1
		}

		if c == '\n' {
			if cf.eof {
				return cfg, nil
			}
			comment = false
			continue
		}
		if comment || isspace(c) {
			continue
		}
		if c == '#' || c == ';' {
			comment = true
			continue
		}
		if c == '[' {
			name, err = cf.getSectionKey()
			if err != nil {
				return cfg, err
			}
			name += "."
			continue
		}
		if !isalpha(c) {
			return cfg, ErrInvalidKeyChar
		}
		key := name + string(c)
		value, err := cf.getValue(&key)
		if err != nil {
			return cfg, err
		}
		cfg[key] = value
	}
}

func (cf *parser) nextChar() byte {
	if len(cf.bytes) == 0 {
		cf.eof = true
		return byte('\n')
	}
	c := cf.bytes[0]
	if c == '\r' {
		/* DOS like systems */
		if len(cf.bytes) > 1 && cf.bytes[1] == '\n' {
			cf.bytes = cf.bytes[1:]
			c = '\n'
		}
	}
	if c == '\n' {
		cf.linenr++
	}
	if len(cf.bytes) == 0 {
		cf.eof = true
		cf.linenr++
		c = '\n'
	}
	cf.bytes = cf.bytes[1:]
	return c
}

func (cf *parser) getSectionKey() (string, error) {
	name := ""
	for {
		c := cf.nextChar()
		if cf.eof {
			return "", ErrUnexpectedEOF
		}
		if c == ']' {
			return name, nil
		}
		if isspace(c) {
			return cf.getExtendedSectionKey(name, c)
		}
		if !iskeychar(c) && c != '.' {
			return "", ErrInvalidSectionChar
		}
		name += string(lower(c))
	}
}

// config: [BaseSection "ExtendedSection"]
func (cf *parser) getExtendedSectionKey(name string, c byte) (string, error) {
	for {
		if c == '\n' {
			cf.linenr--
			return "", ErrSectionNewLine
		}
		c = cf.nextChar()
		if !isspace(c) {
			break
		}
	}
	if c != '"' {
		return "", ErrMissingStartQuote
	}
	name += "."
	for {
		c = cf.nextChar()
		if c == '\n' {
			cf.linenr--
			return "", ErrSectionNewLine
		}
		if c == '"' {
			break
		}
		if c == '\\' {
			c = cf.nextChar()
			if c == '\n' {
				cf.linenr--
				return "", ErrSectionNewLine
			}
		}
		name += string(c)
	}
	if cf.nextChar() != ']' {
		return "", ErrMissingClosingBracket
	}
	return name, nil
}

func (cf *parser) getValue(name *string) (string, error) {
	var c byte
	var err error
	var value string

	/* Get the full name */
	for {
		c = cf.nextChar()
		if cf.eof {
			break
		}
		if !iskeychar(c) {
			break
		}
		*name += string(lower(c))
	}

	for c == ' ' || c == '\t' {
		c = cf.nextChar()
	}

	if c != '\n' {
		if c != '=' {
			return "", ErrInvalidKeyChar
		}
		value, err = cf.parseValue()
		if err != nil {
			return "", err
		}
	}
	/*
	* We already consumed the \n, but we need linenr to point to
	* the line we just parsed during the call to fn to get
	* accurate line number in error messages.
	 */
	// cf.linenr--
	// ret := fn(name->buf, value, data);
	// if ret >= 0 {
	// 	cf.linenr++
	// }
	return value, err
}

func (cf *parser) parseValue() (string, error) {
	var quote, comment bool
	var space int

	var value string

	// strbuf_reset(&cf->value);
	for {
		c := cf.nextChar()
		if c == '\n' {
			if quote {
				cf.linenr--
				return "", ErrUnfinishedQuote
			}
			return value, nil
		}
		if comment {
			continue
		}
		if isspace(c) && !quote {
			if len(value) > 0 {
				space++
			}
			continue
		}
		if !quote {
			if c == ';' || c == '#' {
				comment = true
				continue
			}
		}
		for space != 0 {
			value += " "
			space--
		}
		if c == '\\' {
			c = cf.nextChar()
			switch c {
			case '\n':
				continue
			case 't':
				c = '\t'
				break
			case 'b':
				c = '\b'
				break
			case 'n':
				c = '\n'
				break
				/* Some characters escape as themselves */
			case '\\':
				break
			case '"':
				break
				/* Reject unknown escape sequences */
			default:
				return "", ErrInvalidEscapeSequence
			}
			value += string(c)
			continue
		}
		if c == '"' {
			quote = !quote
			continue
		}
		value += string(c)
	}
}

func lower(c byte) byte {
	return c | 0x20
}

func isspace(c byte) bool {
	return c == '\t' || c == ' ' || c == '\n' || c == '\v' || c == '\f' || c == '\r'
}

func iskeychar(c byte) bool {
	return isalnum(c) || c == '-'
}

func isalnum(c byte) bool {
	return isalpha(c) || isnum(c)
}

func isalpha(c byte) bool {
	return c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z'
}

func isnum(c byte) bool {
	return c >= '0' && c <= '9'
}

func NewRemoteConfig(url string) (*RemoteConfig, error) {
	reSSH, err := regexp2.Compile(regexpGitSSHURL, 0)
	if err != nil {
		return nil, err
	}
	mSSH, err := reSSH.FindStringMatch(url)
	if err != nil {
		return nil, err
	}
	if mSSH != nil {
		groups := mSSH.Groups()
		splited := strings.Split(groups[2].Capture.String(), "/")
		organization := splited[0]
		repository := splited[1]

		return &RemoteConfig{
			URL:          url,
			Organization: organization,
			Repository:   repository,
		}, nil
	}

	reHTTPS, err := regexp2.Compile(regexpGitHTTPSURL, 0)
	if err != nil {
		return nil, err
	}

	mHTTPS, err := reHTTPS.FindStringMatch(url)
	if mHTTPS == nil || err != nil {
		return nil, errors.Wrapf(err, "Not match URL: %v", url)
	}

	groups := mHTTPS.Groups()
	splited := strings.Split(groups[2].Capture.String(), "/")
	organization := splited[1]
	repository := splited[2]
	return &RemoteConfig{
		URL:          url,
		Organization: organization,
		Repository:   repository,
	}, nil
}
