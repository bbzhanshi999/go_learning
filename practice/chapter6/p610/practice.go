package main

func main() {
	addJpg := MakeSuffixFunc("jpg")
	addXml := MakeSuffixFunc("xml")

	println(addJpg("dilireba"))
	println(addXml("mybatis"))
}

func MakeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		return name + "." + suffix
	}
}
