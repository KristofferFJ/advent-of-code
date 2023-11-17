package utils

class LongUtils {
    companion object {
        fun Long.pow(power: Int): Long {
            return this.toBigInteger().pow(power).toLong()
        }

        fun Long.toBase(base: Int): Long {
            var remaining = this
            var result = ""
            while (remaining > 0) {
                val part = remaining.mod(base)
                result += part
                remaining = remaining.div(base)
            }
            return result.reversed().toLong()
        }
    }
}
