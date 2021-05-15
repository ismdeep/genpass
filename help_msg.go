package main

func HelpMsg() string {
	return `Usage: genpass

    -n <LENGTH>            Generate password which length is LENGTH. Default: 16
    --without-fuzzy        Without fuzzy option; default with fuzzy chars
    -d                     Generate with digital
    -a                     Generate with lower case alphabet
    -A                     Generate with UPPER case alphabet
    --hex                  Generate password which is hex
`
}
