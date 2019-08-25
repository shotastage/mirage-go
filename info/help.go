package info

import "fmt"

func Help() {
	fmt.Println("MIRAGE Go")
	fmt.Println("Copyright (C) 2019 Shota Shimazu All Rights Reserved.")
	fmt.Println()
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("		mg <command> [arguments]")
	fmt.Println()
	fmt.Println("The commands are:")
	fmt.Println()
	fmt.Println("		init, initialize       Initialize MIRAGE project & create Mgfile.")
	fmt.Println("		cf, configure          Configure global settings.")
	fmt.Println("		cr, create             Create a new file with headings.")
	fmt.Println("		i, info                Show information.")
	fmt.Println("		h, help                Show usage or help of MIRAGE Go.")
	fmt.Println("		v, version             Show MIRAGE Go version")

	fmt.Println()

}

func infoHelp() {
	fmt.Println("mg info")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("		mg info <option>")
	fmt.Println()
	fmt.Println("The commands are:")
	fmt.Println()
	fmt.Println("		home                   Show home directory.")
	fmt.Println("		os                     Show OS.")
	fmt.Println("		config-file-path       Show config file path.")
	fmt.Println("		check-compatibility    Check MIRAGE Go compatibility.")

	fmt.Println()

}
