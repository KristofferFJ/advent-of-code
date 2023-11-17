package aoc2022.day15.part1

import kotlin.math.absoluteValue

data class Point(val x: Int, val y: Int)
data class SensorBeacon(val sensor: Point, val beacon: Point) {
    fun getCoveredPointsOnLine(line: Int): Set<Point> {
        val distance = (sensor.x - beacon.x).absoluteValue + (sensor.y - beacon.y).absoluteValue
        val xMax = distance - (sensor.y - line).absoluteValue
        return IntRange(-xMax + sensor.x, xMax + sensor.x).map {
            Point(it, line)
        }.toSet()
    }
}

fun main() {
    val sensorBeacons = INPUT.split("\n").map {
        it.replace("Sensor at x=", "")
            .replace(" y=", "")
            .replace(" closest beacon is at x=", "")
            .replace(" y=", "")
            .split(":")
            .map {
                it.split(",")
            }
    }.map {
        SensorBeacon(
            Point(it[0][0].toInt(), it[0][1].toInt()),
            Point(it[1][0].toInt(), it[1][1].toInt()),
        )
    }
    println(sensorBeacons.flatMap {
        it.getCoveredPointsOnLine(2000000)
    }.toSet()
        .minus(
            sensorBeacons.filter {
                it.beacon.y == 2000000
            }.map {
                it.beacon
            }.toSet()
        ).size
    )
}
