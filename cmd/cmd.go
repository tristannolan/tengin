package cmd

type Controller struct {
	Command string
	Valid   []KeyWord
}

type KeyWord struct {
	Identifier string
	Action     func()
}

type Command []KeyWord

func (c *Controller) Execute() {
}

func (c *Controller) ParseCommandPhrase(phrase string) error {
	//cmd := []KeyWord{}
	//for word := range strings.SplitSeq(phrase, " ") {
	//}
	return nil
}

func (c *Controller) Register(k KeyWord) {
}

// :set tick=1
