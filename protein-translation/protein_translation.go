// Protein package consists of tool to help with protein translation.
package protein

import "errors"

var (
	// Error for when we encounter a  stop codon which terminates the translation.
	ErrStop = errors.New("stop codon found")

	// Error for when we encounter an invalid codon / rna.
	ErrInvalidBase = errors.New("provided base is invalid")
)

// FromCodon takes codon as input and returns there corresponding protein.
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

// FromRNA takes a RNA string and translates it into its corresponding proteins.
func FromRNA(rna string) ([]string, error) {

	ln := len(rna)

	// condition to check if the provided rna value is of correct length.
	if ln%3 != 0 {
		return nil, ErrInvalidBase
	}

	var proteinStr []string

	for i := 0; i < ln; i = i + 3 {
		prot, err := FromCodon(rna[i : i+3])

		// condition to check if we found the stop codon.
		if errors.Is(err, ErrStop) {
			return proteinStr, nil
		}

		// condition for if we find a invalid sequence inside the input rna
		if errors.Is(err, ErrInvalidBase) {
			return []string{}, ErrInvalidBase
		}

		proteinStr = append(proteinStr, prot)

	}

	return proteinStr, nil

}
