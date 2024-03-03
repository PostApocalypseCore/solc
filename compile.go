package solc

import (
	"bytes"
	"encoding/json"
	"os/exec"
)

func CompileInDirWithCertainFiles() {

}

type Input struct {
	Lang     lang           `json:"language"`
	Sources  map[string]Src `json:"sources"`
	Settings *Settings      `json:"settings"`
}

type Src struct {
	Keccak256 string   `json:"keccak256,omitempty"`
	Content   string   `json:"content,omitempty"`
	URLS      []string `json:"urls,omitempty"`
}

type Output struct {
	Errors    []error_                       `json:"errors"`
	Sources   map[string]srcOut              `json:"sources"`
	Contracts map[string]map[string]contract `json:"contracts"`
}

func (c *Compiler) BuildSettings(opts []Option) *Settings {
	return c.buildSettings(opts)
}

func (c *Compiler) CompileSingleFile(in *Input) (*Output, error) {
	// init and return on error
	c.once.Do(c.init)
	if c.err != nil {
		return nil, c.err
	}

	inputBuf := bytes.NewBuffer(nil)
	outputBuf := bytes.NewBuffer(nil)

	// encode input
	if err := json.NewEncoder(inputBuf).Encode(in); err != nil {
		return nil, err
	}

	// run solc
	ex := exec.Command(c.solcAbsPath, "--standard-json")
	ex.Stdin = inputBuf
	ex.Stdout = outputBuf
	if err := ex.Run(); err != nil {
		return nil, err
	}

	// decode output
	var output *Output
	if err := json.NewDecoder(outputBuf).Decode(&output); err != nil {
		return nil, err
	}

	return output, nil
}

func GetContract(result map[string]map[string]contract, contractName string) (*Contract, bool) {
	var res *Contract
	var found bool
	for _, conMap := range result {
		for conName, c := range conMap {
			if conName == contractName {
				con := &Contract{
					Code:       c.EVM.DeployedBytecode.Object,
					DeployCode: c.EVM.Bytecode.Object,
					Metadata:   c.Metadata,
				}
				res = con
				return res, true
			}
		}
	}
	return nil, found
}
