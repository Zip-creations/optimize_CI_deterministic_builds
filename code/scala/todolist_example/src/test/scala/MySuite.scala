import org.junit.jupiter.api.Test
import org.junit.jupiter.api.Assertions._
import todolist.*
import scala.collection.mutable.Map


class TestAddingItems {
  @Test
  def testTodo(): Unit = {
    var testlist = TODOList("sample", Map.empty[Int, TODOItem])
    var a = testlist.addItem(TODOItem("testItem1"))
    assertEquals(testlist.items(1).content, "testItem1")
  }
}

class TestRemovingItems {
  @Test
  def testTodo(): Unit = {
    var testlist = TODOList("sample", Map.empty[Int, TODOItem])
    var a = testlist.addItem(TODOItem("testItem1"))
    testlist.removeItemByID(1)
    assertFalse(testlist.items.contains(1))
  }
}
