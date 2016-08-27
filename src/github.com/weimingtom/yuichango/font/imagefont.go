package yuichango

import (
	. "github.com/weimingtom/yuichango/basic"
	. "github.com/weimingtom/yuichango/platform"
)

const (
	FIXMEDFONT_BMP = "fonts/fixedfont.bmp"
	
	FONT_PNG = "fonts/font.png"
	TINY_PNG = "fonts/tiny.png"
	
	FIXEDFONT_BIG_BMP = "fonts/fixedfont_big.bmp"
	FIXEDFONT_BMP = "fonts/fixedfont.bmp"
	PAPYRUS_32_PNG = "fonts/papyrus_32.png"
	RPGFONT_THINNER_NOSHADOW_PNG = "fonts/rpgfont_thinner_noshadow.png"
	RPGFONT_PNG = "fonts/rpgfont.png"
	RPGFONT2_PNG = "fonts/rpgfont2.png"
	TECHYFONTBIG_PNG = "fonts/techyfontbig.png"
	TECHYFONTBIG2_PNG = "fonts/techyfontbig2.png"
	TECHYFONTBLACK_PNG = "fonts/techyfontblack.png"
	TECHYFONTWHITE_PNG = "fonts/techyfontwhite.png"
	THINSANS_BMP = "fonts/thinsans.bmp"
	
	SPACE_ALPHABET_NUMBER = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	SPACE_ALPHABET_NUMBER_MARK = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.,!?-+/():;%&`'*#=[]\""
)

type ImageFont struct {
    mGlyph [256]Rectangle
    mHeight int
    mGlyphSpacing int
    mRowSpacing int
    mImage *Image
    mFilename string	
}
