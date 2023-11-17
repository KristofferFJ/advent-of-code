package utils

class IntUtils {
    companion object {
        infix fun Int.toward(to: Int): IntProgression {
            val step = if (this > to) -1 else 1
            return IntProgression.fromClosedRange(this, to, step)
        }
    }
}