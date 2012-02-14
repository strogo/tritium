
// Returns the index of the last node scoped
@func index(Node %node) Text

// This is how you select an element to pass it to a function.
@func Node.fetch(Text %selector) Text Text

// This is how you access the current node
@func Node.this() Node Node

// Casts a string to a Position node
@func position(Text %position) Position
@func position() Position

// Search the tree and select all matching nodes
@func Node.select(Text %xpath_selector) Text Node

// Opens up the contents to text modification. All XML will get escaped.
@func Node.text() Text Text

// Move the first node, to the second node.
@func Node.move(Node %what, Node %where, Position %pos) Text Node

// Copies the node and yields to it
@func Node.dup() Node Node

// Open's the node's name the currently select Node's name. Aka, change a "div" to a "span"
@func Node.name() Text Text

// Delete the current node
@func Node.remove() Text Node

// Get the XPath of the current node
@func Node.path() Text Text

// Inject some HTML into the node at the prescribed location
@func Node.inject(Text %html) Text

// Create a new node and insert it at the location specified
@func Node.insert_at(Position %pos, Text %tag_name) Text

// Create a new node and insert it at the location specified
@func Node.inject_at(Position %pos, Text %tag_name) Text
@func Node.inject_at(Text %pos, Text %tag_name) Text


@func Node.set(Text %value) Text