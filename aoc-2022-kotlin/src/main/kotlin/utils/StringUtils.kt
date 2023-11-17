package utils

import java.lang.Exception


class StringUtils {
    companion object {
        fun String.isInt(): Boolean {
            return try {
                this.toInt()
                true
            } catch (e: Exception) {
                false
            }
        }

        fun String.toLists(): List<String> {
            val lists = this.substring(1, this.length - 1)
            val elements = mutableListOf<String>()
            var currentElement = ""
            var leftBraces = 0
            var rightBraces = 0
            lists.forEach {
                if(it == '[') { leftBraces +=1; currentElement += it; return@forEach }
                if(it == ']') {
                    rightBraces +=1
                    currentElement += it
                    if(leftBraces == rightBraces) {
                        elements.add(currentElement)
                        currentElement = ""; leftBraces = 0; rightBraces = 0; return@forEach
                    } else {
                        return@forEach
                    }
                }
                if(it == ',') {
                    if(leftBraces == 0) {
                        if(currentElement != "") elements.add(currentElement)
                        currentElement = ""
                        return@forEach
                    } else {
                        currentElement += it
                        return@forEach
                    }
                }
                currentElement += it
            }
            if(currentElement != "") elements.add(currentElement)
            return elements
        }
    }
}