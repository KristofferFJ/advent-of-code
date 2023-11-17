package aoc2022.day16.part2

import utils.CollectionUtils.Companion.isIn
import utils.CollectionUtils.Companion.removeDuplicates

data class Valve(
    val name: String,
    val flowRate: Int,
    private val connections: List<String>,
) {
    fun getConnections(): List<Valve> {
        return valves.filter { it.name in connections }
    }
}

val valves = createValves()

data class DuoPath(
    val humanPath: Path = Path(),
    val elephantPath: Path = Path(),
    var points: Int = 0,
    var valvesTurned: MutableMap<String, Boolean> = valves.associate { it.name to false }.toMutableMap()
) {
    fun turned(valve: String): Boolean {
        return this.valvesTurned[valve]!! || valves.find(valve).flowRate == 0
    }

    fun turnOn(valve: String, round: Int): DuoPath {
        this.valvesTurned[valve] = true
        this.points += valves.find(valve).flowRate * (26 - round)
        return this
    }

    fun onValves(): List<String> {
        return valvesTurned.filter { it.value }.map { it.key }
    }

    override fun toString(): String {
        return "human=${humanPath.position.name}, elephant=${elephantPath.position.name}, points=$points, " +
                "valvesTurned=${valvesTurned.filter { it.value }.map { it.key }}"
    }
}

data class Path(
    var points: Int = 0,
    val position: Valve = valves.find("AA"),
    var passedSinceLastTurnOn: MutableList<String> = mutableListOf("AA")
) {
    fun connections(): List<Valve> {
        val connections = position.getConnections().filter { !it.name.isIn(passedSinceLastTurnOn) }
        return connections.ifEmpty { listOf(position) }
    }
}

data class Duplicate(
    val humanPosition: String,
    val elephantPosition: String,
    val valvesTurned: List<String>
)

var paths = mutableListOf(DuoPath())
fun main() {
    IntRange(1, 26).forEach { round ->
        println(paths.size)
        val tempPaths = paths.toList()
        tempPaths.forEach {
            createRoutes(it, round)
        }
        paths.removeDuplicates({
            Duplicate(it.humanPath.position.name, it.elephantPath.position.name,
                it.valvesTurned.filter { it.value }.map { it.key })
        }, Comparator { a, b -> a.points.compareTo(b.points) })
        paths.removeDuplicates({ setOf(it.humanPath.position, it.elephantPath.position, it.points) })
        paths.removeOutperformed()
        if(round % 5 == 0) {
            paths = paths.sortedBy { it.points }.reversed().subList(0, 500).toMutableList()
        }
    }
    println(paths.sortedBy { it.points }.last())
    println(paths.maxByOrNull { it.points }!!.points)
}

fun createRoutes(path: DuoPath, round: Int) {
    if (!path.turned(path.humanPath.position.name) && !path.turned(path.elephantPath.position.name)
        && path.humanPath.position != path.elephantPath.position
    ) {
        val newPath = path.copy(
            valvesTurned = path.valvesTurned.toMutableMap()
        )
        newPath.turnOn(path.humanPath.position.name, round)
        path.humanPath.passedSinceLastTurnOn = mutableListOf(path.humanPath.position.name)
        newPath.turnOn(path.elephantPath.position.name, round)
        path.elephantPath.passedSinceLastTurnOn = mutableListOf(path.elephantPath.position.name)
        paths.add(newPath)
    }
    if (!path.turned(path.humanPath.position.name)) {
        path.elephantPath.connections().forEach {
            val newPath = path.copy(
                elephantPath = path.elephantPath.copy(position = it),
                valvesTurned = path.valvesTurned.toMutableMap()
            )
            newPath.turnOn(path.humanPath.position.name, round)
            path.humanPath.passedSinceLastTurnOn = mutableListOf(path.humanPath.position.name)
            paths.add(newPath)
        }
    }
    if (!path.turned(path.elephantPath.position.name)) {
        path.humanPath.connections().forEach {
            val newPath = path.copy(
                humanPath = path.humanPath.copy(position = it),
                valvesTurned = path.valvesTurned.toMutableMap()
            )
            newPath.turnOn(path.elephantPath.position.name, round)
            path.elephantPath.passedSinceLastTurnOn = mutableListOf(path.elephantPath.position.name)
            paths.add(newPath)
        }
    }
    path.humanPath.connections().forEach { human ->
        path.elephantPath.connections().forEach { elephant ->
            val newPath = path.copy(
                humanPath = path.humanPath.copy(position = human),
                elephantPath = path.elephantPath.copy(position = elephant),
                valvesTurned = path.valvesTurned.toMutableMap()
            )
            newPath.humanPath.passedSinceLastTurnOn = path.humanPath.passedSinceLastTurnOn.toMutableList()
            newPath.elephantPath.passedSinceLastTurnOn = path.elephantPath.passedSinceLastTurnOn.toMutableList()
            newPath.humanPath.passedSinceLastTurnOn.add(human.name)
            newPath.elephantPath.passedSinceLastTurnOn.add(elephant.name)
            paths.add(newPath)
        }
    }
}

fun MutableList<DuoPath>.removeOutperformed() {
    val grouped = this.groupBy { Pair(it.humanPath.position.name, it.elephantPath.position.name) }
        .filter { it.value.size > 1 }
    val obsoletePaths: List<DuoPath> = grouped.values.map { listWithSamePosition ->
        listWithSamePosition.filter { obsolete ->
            listWithSamePosition.any {
                it.onValves().containsAll(obsolete.onValves()) && it.points > obsolete.points
            }
        }
    }.flatten()
    this.removeAll(obsoletePaths)
}

fun createValves(): List<Valve> {
    return INPUT.split("\n").map {
        it.replace("Valve ", "").replace(" has flow rate", "")
            .replace(" tunnels lead to valves ", "")
            .replace(" tunnel leads to valve ", "")
            .split("=|;".toRegex())
    }.map {
        Valve(it[0], it[1].toInt(), it[2].split(",").map { it.trim() })
    }
}

fun List<Valve>.find(name: String): Valve {
    return this.first { it.name == name }
}
