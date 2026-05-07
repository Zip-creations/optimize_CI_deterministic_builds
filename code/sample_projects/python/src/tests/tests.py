from classes.ToDoList import ToDoList, ToDoItem

def test_add_item():
    todo_list = ToDoList("Alice")
    item1 = ToDoItem("Item 1")
    item2 = ToDoItem("Item 2")
    todo_list.addItem(item1)
    todo_list.addItem(item2)
    print(todo_list.items)
    assert len(todo_list.items) == 2
    assert todo_list.items[0] == item1
    assert todo_list.items[1] == item2
