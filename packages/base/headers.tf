
" Allows you to set global variables. Equivalent to `$`. @example `var(\"txt1\") { set(\"MYTEXT\") }` sets the variable of 'txt1' to 'MYTEXT'."
@func var(Text %name) Text Text
" Allows you to set global variables. Equivalent to `$`. @example `var(\"txt1\", \"MYTEXT\")` sets the variable of 'txt1' to 'MYTEXT'."
@func var(Text %name, Text %value) Text Text

" Returns the time-to-execute (time units vary by implementation). @example `log(time())` will log the time taken to reach that point in the server logs. "
@func time() Text

" Specifies a target (specified by **%target**) to be searched. To be used in conjunction with `with()` - [click for example](http://beta.moovweb.com/learn/training/function_guides/match). @example `match($path) { with(/product/) }` will match the path of the url to see if the regular expression 'product' matches."
@func match(Text %match_target) Text

" Writes out a string (**%log_message**) to the console and debug log. @example `log(\"Importing home-page\")`. "
@func log(Text %log_message) Text Text

" Returns the warning **%message** when a function is deprecated. Mostly useful when defining functions. @example `@func XMLNode.old() { deprecated(\"WARNING! This function has been deprecated\") }` will print out a server log message whenever the function is used. "
@func deprecated(Text %message) Text Text

" Used with `match()` as an opposite of `with()`. @example `match($var) { not('hi') }` will check that the '$var' is not set to 'hi'. "
@func not(Text %text) Text

" Used with `match()` as an opposite of `with()`. @example `match($var) { not(/hi/) }` will check that the '$var' is not set to 'hi'. "
@func not(Regexp %regexp) Text

" Used with `match()`. Allows the match function to specify what is being matched. @example `match($path) { with(\"product\") }` will check the path of the url matches 'product'."
@func with(Text %text) Text

" Used with `match()`. Allows the match function to specify what is being matched. @example `match($path) { with(/product/) }` will check the path of the url matches the regular expression 'product'."
@func with(Regexp %regexp) Text

" Convert from one encoding to another. (If you want a list of encodings, you can run `iconv -l` on your command line.) @example `text() { convert_encoding(\"\", \"\") }` will convert the text from gbk to utf-8."
@func Text.convert_encoding(Text %from, Text %to) Text

" Guess the encoding from the input, the response header and html meta tag. "
@func Text.guess_encoding() Text

" Returns the length of the **%input**. @example `log(length(\"text\"))` will return '4' in the server logs."
@func length(Text %input) Text

" Completes the pseudo-logic of `with()`, allowing the specification of an alternative. @example `match($path) { with(/product/) else() { log(\"Not selecting product\")} }` will only log the message if 'product' is not in the path."
@func else() Text

" Only used within functions - enables functions within the scope of the current function to be performed. @example To learn more, check out [our helpdesk post on how yield works](http://help.moovweb.com/entries/21633781-what-does-the-yield-function-do)."
@func yield() Text

" Only used within functions - enables functions within the scope of the current function to be performed.  @example To learn more, check out [our helpdesk post on how yield works](http://help.moovweb.com/entries/21633781-what-does-the-yield-function-do). "
@func Text.yield() Text

" Returns the current text value. @example `text() { log(this()) }` will log all the text in the DOM into the server logs. "
@func Text.this() Text

" Parses regular expressions. (Use hard-coded regex if you can. This is much slower than hard-coding regex.) The **%options** text allows [Ruby modifiers](http://www.regular-expressions.info/ruby.html) to be included. @example `regexp(\"dog\", \"i\")` uses the `/i` modifier to make the expression case-insensitive, so the regex will match both 'DOG' and 'dog' (plus 'Dog', 'dOg', etc.). "
@func regexp(Text %expression, Text %options) Regexp Text

" Concatenates two (or more) strings. - [click for example](http://beta.moovweb.com/learn/training/function_guides/concat)@example `concat(\"dog\", \"cat\")` is equivalent to `\"dog\" + \"cat\"`. "
@func concat(Text %a, Text %b) Text Text

" This is a the way that we have Tritium communicate variables back to its execution environment. @example `export(\"Content-Type\", \"application/js\")` changes the content-type. "
@func export(Text %key_name) Text Text

" Returns **%input_string** in all caps. @example `upcase(\"dog\")` will return 'DOG'. "
@func upcase(Text %input_string) Text

" Returns **%input_string** in lower case. @example `downcase(\"DOG\")` will return 'dog'. "
@func downcase(Text %input_string) Text

//" Returns the current text scope as a string. Useful to pass the current Text as an argument. "
//@func Text.text() Text Text

" Replaces the entire current text with what you pass in. @example `set(\"one\")` will set the whole of the text to 'one'. "
@func Text.set(Text %value) Text

" Replaces all instances of the regular expression **%search**. This yields to a Text scope that allows you to set the replacement string using `with()` - [click for example](http://beta.moovweb.com/learn/training/function_guides/replace). @example `replace(/bad/) { set(\"good\") }`."
@func Text.replace(Regexp %search) Text Text

" Replaces all instances of the text **%search**. This yields to a Text scope that allows you to set the replacement string using `with()` - [click for example](http://beta.moovweb.com/learn/training/function_guides/replace). @example `replace(\"bad\") { set(\"good\") }`."
@func Text.replace(Text %search) Text Text

" Adds **%text_to_prepend** to the beginning of the text. @example Given `<div>Dog</div>`, `$(\"./div\") { text() { prepend(\"Super-\") } }` will change the text to 'Super-Dog'."
@func Text.prepend(Text %text_to_prepend) Text

" Adds **%text_to_append** to the end of the text. @example Given `<div>Dog</div>`, `$(\"./div\") { text() { append(\"Fish\") } }` will change the text to 'DogFish'."
@func Text.append(Text %text_to_append) Text

" Captures all instances of the regular expression **%search**. "
@func Text.capture(Regexp %search) Text Text

" Rewrite a host/referer/origin from proxy to upstream "
@func Text.rewrite_to_upstream(Text %from_proxy, Text %secure, Text %catchall) Text

" Rewrite a link from upstream to proxy "
@func Text.rewrite_to_proxy(Text %secure, Text %catchall) Text

" Rewrite a cookie domain from upstream to proxy "
@func Text.rewrite_cookie_domain(Text %host, Text %secure, Text %catchall) Text

" Rewrite a link from upstream to proxy "
@func Text.rewrite_link(Text %secure, Text %catchall) Text
