package aoc2022.day18.part2

private val cubes = INPUT.split("\n").map { it.split(",") }.map {
    Cube(
        it[0].toInt(),
        it[1].toInt(),
        it[2].toInt()
    )
}

val maxX = cubes.maxBy { it.x }.x
val maxY = cubes.maxBy { it.y }.y
val maxZ = cubes.maxBy { it.z }.z
private val goodPaths = mutableSetOf<Cube>()
private val badPaths = mutableSetOf<Cube>()

private data class Point(val x: Int, val y: Int, val z: Int)
private data class Side(val vectorStart: Point, val vectorEnd: Point) {
    fun isFacingOut(): Boolean {
        val cube = listOf(
            Cube(vectorStart.x, vectorStart.y, vectorStart.z),
            Cube(
                if (vectorEnd.x == vectorStart.x) vectorStart.x - 1 else vectorStart.x,
                if (vectorEnd.y == vectorStart.y) vectorStart.y - 1 else vectorStart.y,
                if (vectorEnd.z == vectorStart.z) vectorStart.z - 1 else vectorStart.z,
            )
        ).first { !cubes.contains(it) }
        val isFacingOut = cube.pathLeadsOut(mutableListOf(cube), mutableListOf())
        if (!isFacingOut) {
            badPaths.add(cube)
        }
        return isFacingOut
    }
}

private data class Cube(val x: Int, val y: Int, val z: Int) {
    fun getSides(): Set<Side> {
        return setOf(
            Side(Point(x, y, z), Point(x + 1, y + 1, z)),
            Side(Point(x, y, z), Point(x + 1, y, z + 1)),
            Side(Point(x, y, z), Point(x, y + 1, z + 1)),
            Side(Point(x + 1, y, z), Point(x + 1, y + 1, z + 1)),
            Side(Point(x, y + 1, z), Point(x + 1, y + 1, z + 1)),
            Side(Point(x, y, z + 1), Point(x + 1, y + 1, z + 1)),
        )
    }

    fun getConnectingCubes(): List<Cube> {
        return listOf(
            Cube(x + 1, y, z),
            Cube(x, y + 1, z),
            Cube(x, y, z + 1),
            Cube(x - 1, y, z),
            Cube(x, y - 1, z),
            Cube(x, y, z - 1),
        )
    }

    fun pathLeadsOut(path: MutableList<Cube>, deadEnds: MutableList<Cube>): Boolean {
        if (this.x > maxX || this.y > maxY || this.z > maxZ) {
            return true
        }
        if (path.intersect(badPaths).isNotEmpty()) {
            badPaths.addAll(path)
            return false
        }
        if (path.intersect(goodPaths).isNotEmpty()) {
            goodPaths.addAll(path)
            return true
        }
        val pathLeadsOut = this.getConnectingCubes()
            .filter { !cubes.contains(it) }
            .filter { !path.contains(it) }
            .filter { !deadEnds.contains(it) }
            .any { it.pathLeadsOut((path + it).toMutableList(), deadEnds) }
        if (pathLeadsOut) {
            goodPaths.addAll(path)
            return true
        }
        deadEnds.add(path.last())
        return false
    }
}

fun main() {
    val sides = cubes.flatMap { it.getSides() }
    val surfaceSides = sides.groupBy { it }
        .filter { it.value.size == 1 }
        .keys
    println(surfaceSides.filter { it.isFacingOut() }.size)
}

