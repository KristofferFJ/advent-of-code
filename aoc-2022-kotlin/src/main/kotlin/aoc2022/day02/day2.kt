package aoc2022.day02

fun getRounds(): List<List<String>> {
    return INPUT.split("\n").map { it.split(" ") }
}

// X og A er rock, 1 point
// Y og B er paper, 2 point
// Z og C er scissor, 3 point

fun calculateRoundOne(round: List<String>): Int {
    if(round[1] == "X") {
        if(round[0] == "A") return 4
        if(round[0] == "B") return 1
        if(round[0] == "C") return 7
    }
    if(round[1] == "Y") {
        if(round[0] == "A") return 8
        if(round[0] == "B") return 5
        if(round[0] == "C") return 2
    }
    if(round[1] == "Z") {
        if(round[0] == "A") return 3
        if(round[0] == "B") return 9
        if(round[0] == "C") return 6
    }
    throw Exception("")
}

//X lose, Y draw, Z win
fun calculateRoundTwo(round: List<String>): Int {
    if(round[1] == "X") {
        if(round[0] == "A") return 0 + 3
        if(round[0] == "B") return 0 + 1
        if(round[0] == "C") return 0 + 2
    }
    if(round[1] == "Y") {
        if(round[0] == "A") return 3 + 1
        if(round[0] == "B") return 3 + 2
        if(round[0] == "C") return 3 + 3
    }
    if(round[1] == "Z") {
        if(round[0] == "A") return 6 + 2
        if(round[0] == "B") return 6 + 3
        if(round[0] == "C") return 6 + 1
    }
    throw Exception("")
}

fun partOne() {
    println(getRounds().sumOf { calculateRoundOne(it) })
}

fun partTwo() {
    println(getRounds().sumOf { calculateRoundTwo(it) })
}

fun main() {
    partOne()
    partTwo()
}