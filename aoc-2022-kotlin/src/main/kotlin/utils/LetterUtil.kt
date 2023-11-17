package utils

class LetterUtil {
    companion object {
        private val LETTERS = listOf(
            'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
            'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
        )

        fun Char.getLetterValue(): Int {
            if (this.isUpperCase()) {
                return LETTERS.indexOf(this.lowercaseChar()) + 27
            }
            return LETTERS.indexOf(this) + 1
        }
    }
}
