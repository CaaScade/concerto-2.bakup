package manager

func Compile(files []string) error {
	if err := PreProcess(files); err != nil {
		return err
	}

	trees, err := Parse(files)
	if err != nil {
		return err
	}

	symTab, err := ExtractSymbolTable(trees)
	if err != nil {
		return err
	}

	_ = symTab

	return nil
}
