package test

import java.io.File

fun readLine(): List<String> {
    return File("src/main/kotlin/test/input.txt").readLines()
}
