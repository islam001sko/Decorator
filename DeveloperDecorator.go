package main

type DeveloperDecorator struct {
	developer Developer
}

func (d *DeveloperDecorator) writeCode() string {
	return d.writeCode()
}
