package todolist
import scala.collection.mutable.Map

class TODOList(
  var author: String,
  var items: Map[Int, TODOItem]
) {
  var id = 0
  def newID = () =>
    id += 1
    id

  def addItem = (item: TODOItem) =>
    items(newID()) = item
  
  def removeItemByID = (id: Int) =>
    items -= id
  
  def removeAllItems = () =>
    items.clear()
  
  def removeDuplicateItems = () =>
    var seen = Set.empty[String]
    items = items.filter { case (id, item) =>
      if (seen.contains(item.content)) {
        false
      } else {
        seen += item.content
        true
      }
    }

  def getAuthor = () =>
    author
  
  def getItems = () =>
    items.values.toList
}

case class TODOItem (
  content: String
)
