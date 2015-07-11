package pcre

// Reference to compile pattern
type Regexp struct {
	pattern string
}

// Helper function to get quick answer
func QuickMatch(pattern string, subject string, flags int) (bool, error) {
	re, err := Compile(pattern, flags)
	if err != nil {
		return false, err
	}
	m := re.Matcher(subject, flags)
	return m.Match(subject, flags), nil
}

// Compiles pattern with flags, error is nil on success
func Compile(pattern string, flags int) (Regexp, error) {
	return Regexp{}, nil
}

// Returns number of groups for compiled pattern
func (re *Regexp) Groups() int {
	return -1
}

//Reference to matched results
type Matcher struct {
	re      Regexp
	groups  int
	subject string
	matches bool
}

//Return matcher for subject
func (re *Regexp) Matcher(subject string, flags int) (m *Matcher) {
	return &Matcher{}
}

//Reset match with new Regex, subject, flags
func (m *Matcher) Reset(re Regexp, subject string, flags int) {
	//
}

//Returns whether subject matched Matchers
func (m *Matcher) Match(subject string, flags int) bool {
	return m.matches
}

//Number of groups in Matchee
func (m *Matcher) Groups() int {
	return m.groups
}

//Return named group (?)
func (m *Matcher) Named(group string) string {
	return ""
}
