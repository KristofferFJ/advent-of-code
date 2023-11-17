package aoc2022.day07

fun main() {
    fun createDirectory(): Directory {
        val commands = INPUT.split("\n")
        val outerDirectory = Directory("/", mutableListOf(), mutableListOf(), null)
        var currentDirectory = outerDirectory
        for (i in 1 until commands.size) {
            if (commands[i] == "$ ls") {
                continue
            } else if (commands[i].startsWith("dir")) {
                continue
            } else if (commands[i].matches("^\\d+.*$".toRegex())) {
                val (size, name) = commands[i].split(" ")
                currentDirectory.files.add(File(name, size.toLong()))
            } else if (commands[i] == "$ cd ..") {
                currentDirectory = currentDirectory.outerDirectory!!
            } else if (commands[i].startsWith("$ cd")) {
                val (_, _, name) = commands[i].split(" ")
                currentDirectory = if (currentDirectory.directories.find { it.name == name } != null) {
                    currentDirectory.directories.find { it.name == name }!!
                } else {
                    val newDirectory = Directory(name, mutableListOf(), mutableListOf(), currentDirectory)
                    currentDirectory.directories.add(
                        newDirectory
                    )
                    newDirectory
                }
            }
        }
        return outerDirectory
    }

    fun partOne() {
        val flatDirectory = mutableListOf<Directory>()
        createDirectory().flatDirectories(flatDirectory)
        println(flatDirectory.filter { it.size() < 100000 }.sumOf { it.size() })
    }

    fun partTwo() {
        val totalSize = 70000000
        val requiredSize = 30000000
        val directories = createDirectory()
        val missingSize = requiredSize - (totalSize - directories.size())
        val flatDirectory = mutableListOf<Directory>()
        println(directories.flatDirectories(flatDirectory).filter { it.size() > missingSize }.minBy { it.size() }.size())
    }

    partOne()
    partTwo()
}

data class File(val name: String, val size: Long)
data class Directory(
    val name: String,
    val files: MutableList<File>,
    val directories: MutableList<Directory>,
    val outerDirectory: Directory?
) {
    fun size(): Long {
        return this.files.sumOf { it.size } + this.directories.sumOf { it.size() }
    }

    fun flatDirectories(flatDirectory: MutableList<Directory>): MutableList<Directory> {
        flatDirectory.add(this)
        directories.forEach { it.flatDirectories(flatDirectory) }
        return flatDirectory
    }
}
