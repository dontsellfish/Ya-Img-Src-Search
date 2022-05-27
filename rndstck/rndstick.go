package rndstck

import (
	"gopkg.in/telebot.v3"
	"math/rand"
	"time"
)

var StickerPool = []string{
	"CAACAgIAAxkBAAEEG-5iKuHxy4xE045GhwOfkdW2y9xvlAACaQAD_DykE9Y6t3D0xHBIIwQ",
	"CAACAgIAAxkBAAEEG_BiKuHzSaEXh6FUZD0QvZN0Z1YqvgACdQAD_DykEyblAjF1Ko45IwQ",
	"CAACAgIAAxkBAAEEG_JiKuICi1z24905hguXGRj4xj8lpAACegAD_DykE-on7yosaJYiIwQ",
	"CAACAgIAAxkBAAEEG_RiKuIO1x3GQQ5715tuPBd4nPbOcwACagAD_DykE4fhjBwS2G5rIwQ",
	"CAACAgIAAxkBAAEEG_ZiKuIPxx0YLFCi4CPuuH53FhYv2gACawAD_DykE7Q45_7__hUZIwQ",
	"CAACAgIAAxkBAAEEG_hiKuIWJMuOb0NTfZclk5EoaQABDW8AAmwAA_w8pBNuCyJY53td_yME",
	"CAACAgIAAxkBAAEEG_piKuIXAZWPLHrCdRdCCQABsSnzaQgAAm0AA_w8pBPCKLn5q8ls0SME",
	"CAACAgIAAxkBAAEEG_xiKuIZ6arXgMyRl3yLSSxDBwPliwACiAAD_DykE6ZhRgJgXUDDIwQ",
	"CAACAgIAAxkBAAEEG_5iKuIbriOGPXYs4kH-kBG-wpGzDwACewAD_DykE864AAHndwzFICME",
	"CAACAgIAAxkBAAEEHAJiKuIn760hJZfTbyJnVCaikDAzwAACgwAD_DykEzyJbu3q5wr6IwQ",
	"CAACAgIAAxkBAAEEHARiKuIqOC_h544nnEuHsTdth--IDAACcQAD_DykE9zdEEKY6fjrIwQ",
	"CAACAgIAAxkBAAEEHAZiKuIs6A57jdNRnGsyZDNl2SXJJAACcgAD_DykE2lhXtpeCIJHIwQ",
	"CAACAgIAAxkBAAEEHAhiKuIumM0diMZ5R7U1DlVIo4WUlQACcwAD_DykEyDXzjgL7IqDIwQ",
	"CAACAgIAAxkBAAEEHApiKuIvkCDT_3qTVmEp6G0P5jVUYQACmgAD_DykE2AzvbVqb20mIwQ",
	"CAACAgIAAxkBAAEEHAxiKuIx_L_G8ZeMw_upIuj0e0_fAwACbgAD_DykEyAGiVayl7OQIwQ",
	"CAACAgIAAxkBAAEEHA5iKuI0hoWilmG3J6Y75E9KhMi4pwACbwAD_DykE0jDtCTZG84MIwQ",
	"CAACAgIAAxkBAAEEHBBiKuI20l4RMDv5e91WDu6Bi3dRjQACcAAD_DykE0iNOE4vyF12IwQ",
	"CAACAgIAAxkBAAEEHBJiKuI4eUatxh-nNaZecMukH5iYDQACdgAD_DykEwMmjLx6kSBVIwQ",
	"CAACAgIAAxkBAAEEHBRiKuI5mXfqudW0-tHX9IMN-mO2xwACdwAD_DykE8X_5byEHBVtIwQ",
	"CAACAgIAAxkBAAEEHBZiKuI7ekK7Q34FmzOQ2fSMdUuFMAACdwAD_DykE8X_5byEHBVtIwQ",
	"CAACAgIAAxkBAAEEHBhiKuI9eSiI1QXOVfiFh3AsSH-6xgACeAAD_DykE5hm_lRkZGP5IwQ",
	"CAACAgIAAxkBAAEEHBpiKuI_S9DWNPikD3sWU8Mjw9AqmAACfAAD_DykE4Fv3UK_ilqAIwQ",
	"CAACAgIAAxkBAAEEHBxiKuJBwwwzQ8VrpdRdt1u55l1keQACfgAD_DykE9pDJhJfmU53IwQ",
	"CAACAgIAAxkBAAEEHB5iKuJCZGQvqlSHwm0146m50wxeJwACfwAD_DykExACcDtc9_6yIwQ",
	"CAACAgIAAxkBAAEEHCBiKuJEV_wwPzn7oijblDW8OWqH3AACgAAD_DykE056R_cPz8jeIwQ",
	"CAACAgIAAxkBAAEEHCJiKuJG9sccWOqd1DDFfvvG4DV8_gACgQAD_DykE8-K_Dz2UzuvIwQ",
	"CAACAgIAAxkBAAEEHCRiKuJIaN5UtM3jo7C0XIikdoKo0QACggAD_DykEw_m8de9bxwbIwQ",
	"CAACAgIAAxkBAAEEHCZiKuJKmC3PbJkZPmnY8nKFqtU6mgACiQAD_DykEza-7I7Fr8euIwQ",
	"CAACAgIAAxkBAAEEHChiKuJMNTAj3OhZuiwatLa2MlYX9gACigAD_DykE0yBW8NHCmkFIwQ",
	"CAACAgIAAxkBAAEEHCpiKuJNwqqbl2E2iaj40S2hnoYyDgACiwAD_DykE0obZERT5hb6IwQ",
	"CAACAgIAAxkBAAEEHCxiKuJPAw8SfYAjXjwDgbXshlLwAgACkAAD_DykE2MuOopBhczsIw",
}

var randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

// Get
// There's a chance that changing StickerPool could cause data race and this function panics.
// So, be careful with changing StickerPool
func Get() (sticker *telebot.Sticker) {
	i := randomGenerator.Int() % len(StickerPool)
	stickerID := StickerPool[i]
	return &telebot.Sticker{File: telebot.File{FileID: stickerID}}
}
