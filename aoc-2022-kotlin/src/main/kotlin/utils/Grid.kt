package utils

data class Point(val x: Int, val y: Int) {
    fun addX(value: Int): Point = Point(x + value, y)
    fun addY(value: Int): Point = Point(x, y + value)
}

data class Field(val point: Point, var value: String) {
    val x = point.x
    val y = point.y
}
class Grid(stringInput: String) {
    var rows: MutableList<MutableList<Field>> = stringInput.split("\n").mapIndexed { rowIndex, row ->
        row.toList().mapIndexed { columnIndex, value ->
            Field(Point(columnIndex, rowIndex), value.toString())
        }.toMutableList()
    }.toMutableList()

    fun getFields() = rows.flatten()
    private fun getColumnCount() = rows[0].size
    private fun getRowCount() = rows.size
    private fun getMinX() = rows[0][0].x
    private fun getMinY() = rows[0][0].y
    private fun getMaxX() = rows[getRowCount() - 1][getColumnCount() - 1].point.x
    private fun getMaxY() = rows[getRowCount() - 1][getColumnCount() - 1].point.y

    fun at(point: Point): String {
        return get(point.x, point.y).value
    }

    fun get(x: Int, y: Int): Field {
        expandGridIfNecessary(x, y)
        return rows.find { it[0].y == y }!!.find { it.x == x }!!
    }

    fun set(x: Int, y: Int, value: String) {
        expandGridIfNecessary(x, y)
        rows.find { it[0].y == y }!!.find { it.x == x }!!.value = value
    }

    private fun expandGridIfNecessary(x: Int, y: Int) {
        if (x > getMaxX() || x < getMinX()) {
            addColumns(x)
        }
        if (y > getMaxY() || y < getMinY()) {
            addRows(y)
        }
    }

    private fun addColumns(x: Int) {
        if (x > getMaxX()) {
            val range = IntRange(getMaxX() + 1, x)
            rows.forEach { row ->
                row.addAll(range.map { Field(Point(it, row[0].point.y), ".") })
            }
        } else {
            val range = IntRange(x, getMinX() - 1)
            rows.forEach { row ->
                row.addAll(0, range.map { Field(Point(it, row[0].point.y), ".") })
            }
        }
    }

    private fun addRows(y: Int) {
        if (y > getMaxY()) {
            val range = IntRange(getMinX(), getMaxX())
            IntRange(getMaxY() + 1, y).forEach { row ->
                rows.add(range.map { Field(Point(it, row), ".") }.toMutableList())
            }
        } else {
            val range = IntRange(getMinX(), getMaxX())
            IntRange(getMinY() - 1, y).forEach { row ->
                rows.add(0, range.map { Field(Point(it, row), ".") }.toMutableList())
            }
        }
    }

    override fun toString(): String {
        return rows.reversed().joinToString("") {
            "${it.joinToString("") { it.value }}\n"
        }
    }
}
