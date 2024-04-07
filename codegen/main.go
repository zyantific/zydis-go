package main

import (
	"log"

	"github.com/can1357/gengo/clang"
	"github.com/can1357/gengo/gengo"
)

type zydisProvider struct {
	*gengo.BaseProvider
}

func (p *zydisProvider) NameField(name string, recordName string) string {
	if recordName == "ZydisDecodedInstructionRawEvex_" {
		if name == "b" {
			return "Br"
		}
	}
	return p.BaseProvider.NameField(name, recordName)
}

func main() {
	prov := &zydisProvider{
		BaseProvider: gengo.NewBaseProvider(
			gengo.WithRemovePrefix(
				"Zydis_", "Zyan_", "Zycore_",
				"Zydis", "Zyan", "Zycore",
			),
			gengo.WithInferredMethods([]gengo.MethodInferenceRule{
				{Name: "ZydisDecoder", Receiver: "Decoder"},
				{Name: "ZydisEncoder", Receiver: "EncoderRequest"},
				{Name: "ZydisFormatterBuffer", Receiver: "FormatterBuffer"},
				{Name: "ZydisFormatter", Receiver: "ZydisFormatter *"},
				{Name: "ZyanVector", Receiver: "Vector"},
				{Name: "ZyanStringView", Receiver: "StringView"},
				{Name: "ZyanString", Receiver: "String"},
				{Name: "ZydisRegister", Receiver: "Register"},
				{Name: "ZydisMnemonic", Receiver: "Mnemonic"},
				{Name: "ZydisISASet", Receiver: "ISASet"},
				{Name: "ZydisISAExt", Receiver: "ISAExt"},
				{Name: "ZydisCategory", Receiver: "Category"},
			}),
			gengo.WithForcedSynthetic(
				"ZydisShortString_",
				"struct ZydisShortString_",
			),
		),
	}
	pkg := gengo.NewPackageWithProvider("zydis", prov)
	err := pkg.Transform("zydis", &clang.Options{
		Sources: []string{"./Zydis.h"},
		AdditionalParams: []string{
			"-DZYAN_NO_LIBC",
			"-DZYAN_STATIC_ASSERT",
		},
	})
	if err != nil {
		log.Fatalf("Failed to transform: %v", err)
	}

	if err := pkg.WriteToDir("../"); err != nil {
		log.Fatalf("Failed to write the directory: %v", err)
	}
}
