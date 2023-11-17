package aoc2022.day15.part2

import kotlin.math.absoluteValue

const val maxSize = 4000000

data class Point(val x: Int, val y: Int)
data class SensorBeacon(val sensor: Point, val beacon: Point) {
    private val viewingDistance = (sensor.x - beacon.x).absoluteValue + (sensor.y - beacon.y).absoluteValue

    fun getCoveredPointsOnLine(line: Int): Interval? {
        val viewingDistance = viewingDistance
        val distanceToLine = (sensor.y - line).absoluteValue
        val maxSpotsToCover = viewingDistance - (sensor.y - line).absoluteValue
        if (maxSpotsToCover < 1) return null
        val from = maxOf(sensor.x - (viewingDistance - distanceToLine), -1)
        val to = minOf(sensor.x + (viewingDistance - distanceToLine), maxSize + 1)
        return Interval(minOf(from, maxSize + 1), maxOf(to, -1))
    }
}

data class Interval(var start: Int, var end: Int) {
    fun intersects(interval: Interval): Boolean {
        return interval.start <= this.end + 1 && this.start <= interval.end + 1
    }
}
data class CoveredIntervals(val intervals: MutableList<Interval> = mutableListOf()) {
    fun addIntervalAndCheckOverlap(interval: Interval) {
        val intersectingInterval = this.intervals.find { it.intersects(interval) }
        if(intersectingInterval != null) {
            intersectingInterval.start = minOf(interval.start, intersectingInterval.start)
            intersectingInterval.end = maxOf(interval.end, intersectingInterval.end)
            spliceIntervals(intersectingInterval)
            return
        }
        this.intervals.add(interval)
    }

    private fun spliceIntervals(newInterval: Interval) {
        val intersectingInterval = this.intervals.find { it.intersects(newInterval) && it != newInterval }
        if(intersectingInterval != null) {
            intersectingInterval.start = minOf(newInterval.start, intersectingInterval.start)
            intersectingInterval.end = maxOf(newInterval.end, intersectingInterval.end)
            this.intervals.remove(newInterval)
            return spliceIntervals(intersectingInterval)
        }
    }

    fun isComplete(): Boolean {
        return this.intervals.size == 1 && this.intervals[0].start == -1 && this.intervals[0].end == maxSize + 1
    }
}

fun getAvailablePointsOnLine(sensorBeacons: List<SensorBeacon>, y: Int) {
    val intervals = CoveredIntervals()
    sensorBeacons.forEach {
        val coveredPoints = it.getCoveredPointsOnLine(y) ?: return@forEach
        intervals.addIntervalAndCheckOverlap(coveredPoints)
        if(intervals.isComplete()) {
            return
        }
    }
    println(
        (intervals.intervals[0].end + 1).toLong() * 4000000L + y
    )
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
    IntRange(0, maxSize).forEach {
        getAvailablePointsOnLine(sensorBeacons, it)
    }
}
