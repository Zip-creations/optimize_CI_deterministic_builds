import scala.collection.mutable.Map
import todolist.*

@main def hello(): Unit =
  println("Hello world!")
  println(msg)
  var testlist = TODOList("sample", Map.empty[Int, TODOItem])
  var a = testlist.addItem(TODOItem("testItem1"))
  print(testlist.items.toString + "\n")
  testlist.removeItemByID(1)
  print(testlist.items.toString + "\n")

def msg = "I was compiled by Scala 3. :)"
