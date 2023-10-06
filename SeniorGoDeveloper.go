package main

type SeniorGoDeveloper struct {
	d Developer
}

func (s *SeniorGoDeveloper) makeCodeReview() string {
	return "Code Review"
}

func (s *SeniorGoDeveloper) writeCode() string {
	return s.d.writeCode() + s.makeCodeReview()
}
