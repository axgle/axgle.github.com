<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/datafmt/datafmt.go</title>

  <link rel="stylesheet" type="text/css" href="../../../../doc/style.css">
  <script type="text/javascript" src="../../../../doc/godocs.js"></script>

</head>

<body>

  <script>
    // Catch 'enter' key down events and trigger the search form submission.
    function codesearchKeyDown(event) {
      if (event.which == 13) {
        var form = document.getElementById('codesearch');
        var query = document.getElementById('codesearchQuery');
        form.q.value = "lang:go package:go.googlecode.com " + query.value;
        document.getElementById('codesearch').submit();
}      return true;
}
    // Capture the submission event and construct the query parameter.
    function codeSearchSubmit() {
      var query = document.getElementById('codesearchQuery');
      var form = document.getElementById('codesearch');
      form.q.value = "lang:go package:go.googlecode.com " + query.value;
      return true;
}  </script>

<div id="topnav">
  <table summary="">
    <tr>
      <td id="headerImage">
        <a href="../../../../index.html"><img src="../../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
      </td>
      <td>
        <div id="headerDocSetTitle">The Go Programming Language</div>
      </td>
      <td>
        <!-- <table>
          <tr>
            <td>
              <! The input box is outside of the form because we want to add
              a couple of restricts to the query before submitting. If we just
              add the restricts to the text box before submitting, then they
              appear in the box when the user presses 'back'. Thus we use a
              hidden field in the form. However, there's no way to stop the
              non-hidden text box from also submitting a value unless we move
              it outside of the form
              <input type="search" id="codesearchQuery" value="" size="30" onkeydown="return codesearchKeyDown(event);"/>
              <form method="GET" action="http://www.google.com/codesearch" id="codesearch" class="search" onsubmit="return codeSearchSubmit();" style="display:inline;">
                <input type="hidden" name="q" value=""/>
                <input type="submit" value="Code search" />
                <span style="color: red">(TODO: remove for now?)</span>
              </form>
            </td>
          </tr>
          <tr>
            <td>
              <span style="color: gray;">(e.g. &ldquo;pem&rdquo; or &ldquo;xml&rdquo;)</span>
            </td>
          </tr>
        </table> -->
      </td>
    </tr>
  </table>
</div>

<div id="linkList">
  <ul>
    <li class="navhead"><a href="../../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../../index.html">Source files</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Help</li>
    <li>#go-nuts on irc.freenode.net</li>
    <li><a href="http://groups.google.com/group/golang-nuts">Go Nuts mailing list</a></li>
    <li><a href="http://code.google.com/p/go/issues/list">Issue tracker</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Go code search</li>
    <form method="GET" action="http://golang.org/search" class="search">
    <input type="search" name="q" value="" size="25" style="width:80%; max-width:200px" />
    <input type="submit" value="Go" />
    </form>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Last update</li>
	<li>Thu Nov 12 15:46:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/datafmt/datafmt.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">/*	The datafmt package implements syntax-directed, type-driven formatting</span>
<a id="L6"></a><span class="comment">	of arbitrary data structures. Formatting a data structure consists of</span>
<a id="L7"></a><span class="comment">	two phases: first, a parser reads a format specification and builds a</span>
<a id="L8"></a><span class="comment">	&#34;compiled&#34; format. Then, the format can be applied repeatedly to</span>
<a id="L9"></a><span class="comment">	arbitrary values. Applying a format to a value evaluates to a []byte</span>
<a id="L10"></a><span class="comment">	containing the formatted value bytes, or nil.</span>

<a id="L12"></a><span class="comment">	A format specification is a set of package declarations and format rules:</span>

<a id="L14"></a><span class="comment">		Format      = [ Entry { &#34;;&#34; Entry } [ &#34;;&#34; ] ] .</span>
<a id="L15"></a><span class="comment">		Entry       = PackageDecl | FormatRule .</span>

<a id="L17"></a><span class="comment">	(The syntax of a format specification is presented in the same EBNF</span>
<a id="L18"></a><span class="comment">	notation as used in the Go language specification. The syntax of white</span>
<a id="L19"></a><span class="comment">	space, comments, identifiers, and string literals is the same as in Go.)</span>

<a id="L21"></a><span class="comment">	A package declaration binds a package name (such as &#39;ast&#39;) to a</span>
<a id="L22"></a><span class="comment">	package import path (such as &#39;&#34;go/ast&#34;&#39;). Each package used (in</span>
<a id="L23"></a><span class="comment">	a type name, see below) must be declared once before use.</span>

<a id="L25"></a><span class="comment">		PackageDecl = PackageName ImportPath .</span>
<a id="L26"></a><span class="comment">		PackageName = identifier .</span>
<a id="L27"></a><span class="comment">		ImportPath  = string .</span>

<a id="L29"></a><span class="comment">	A format rule binds a rule name to a format expression. A rule name</span>
<a id="L30"></a><span class="comment">	may be a type name or one of the special names &#39;default&#39; or &#39;/&#39;.</span>
<a id="L31"></a><span class="comment">	A type name may be the name of a predeclared type (for example, &#39;int&#39;,</span>
<a id="L32"></a><span class="comment">	&#39;float32&#39;, etc.), the package-qualified name of a user-defined type</span>
<a id="L33"></a><span class="comment">	(for example, &#39;ast.MapType&#39;), or an identifier indicating the structure</span>
<a id="L34"></a><span class="comment">	of unnamed composite types (&#39;array&#39;, &#39;chan&#39;, &#39;func&#39;, &#39;interface&#39;, &#39;map&#39;,</span>
<a id="L35"></a><span class="comment">	or &#39;ptr&#39;). Each rule must have a unique name; rules can be declared in</span>
<a id="L36"></a><span class="comment">	any order.</span>

<a id="L38"></a><span class="comment">		FormatRule  = RuleName &#34;=&#34; Expression .</span>
<a id="L39"></a><span class="comment">		RuleName    = TypeName | &#34;default&#34; | &#34;/&#34; .</span>
<a id="L40"></a><span class="comment">		TypeName    = [ PackageName &#34;.&#34; ] identifier .</span>

<a id="L42"></a><span class="comment">	To format a value, the value&#39;s type name is used to select the format rule</span>
<a id="L43"></a><span class="comment">	(there is an override mechanism, see below). The format expression of the</span>
<a id="L44"></a><span class="comment">	selected rule specifies how the value is formatted. Each format expression,</span>
<a id="L45"></a><span class="comment">	when applied to a value, evaluates to a byte sequence or nil.</span>

<a id="L47"></a><span class="comment">	In its most general form, a format expression is a list of alternatives,</span>
<a id="L48"></a><span class="comment">	each of which is a sequence of operands:</span>

<a id="L50"></a><span class="comment">		Expression  = [ Sequence ] { &#34;|&#34; [ Sequence ] } .</span>
<a id="L51"></a><span class="comment">		Sequence    = Operand { Operand } .</span>

<a id="L53"></a><span class="comment">	The formatted result produced by an expression is the result of the first</span>
<a id="L54"></a><span class="comment">	alternative sequence that evaluates to a non-nil result; if there is no</span>
<a id="L55"></a><span class="comment">	such alternative, the expression evaluates to nil. The result produced by</span>
<a id="L56"></a><span class="comment">	an operand sequence is the concatenation of the results of its operands.</span>
<a id="L57"></a><span class="comment">	If any operand in the sequence evaluates to nil, the entire sequence</span>
<a id="L58"></a><span class="comment">	evaluates to nil.</span>

<a id="L60"></a><span class="comment">	There are five kinds of operands:</span>

<a id="L62"></a><span class="comment">		Operand     = Literal | Field | Group | Option | Repetition .</span>

<a id="L64"></a><span class="comment">	Literals evaluate to themselves, with two substitutions. First,</span>
<a id="L65"></a><span class="comment">	%-formats expand in the manner of fmt.Printf, with the current value</span>
<a id="L66"></a><span class="comment">	passed as the parameter. Second, the current indentation (see below)</span>
<a id="L67"></a><span class="comment">	is inserted after every newline or form feed character.</span>

<a id="L69"></a><span class="comment">		Literal     = string .</span>

<a id="L71"></a><span class="comment">	This table shows string literals applied to the value 42 and the</span>
<a id="L72"></a><span class="comment">	corresponding formatted result:</span>

<a id="L74"></a><span class="comment">		&#34;foo&#34;       foo</span>
<a id="L75"></a><span class="comment">		&#34;%x&#34;        2a</span>
<a id="L76"></a><span class="comment">		&#34;x = %d&#34;    x = 42</span>
<a id="L77"></a><span class="comment">		&#34;%#x = %d&#34;  0x2a = 42</span>

<a id="L79"></a><span class="comment">	A field operand is a field name optionally followed by an alternate</span>
<a id="L80"></a><span class="comment">	rule name. The field name may be an identifier or one of the special</span>
<a id="L81"></a><span class="comment">	names @ or *.</span>

<a id="L83"></a><span class="comment">		Field       = FieldName [ &#34;:&#34; RuleName ] .</span>
<a id="L84"></a><span class="comment">		FieldName   = identifier | &#34;@&#34; | &#34;*&#34; .</span>

<a id="L86"></a><span class="comment">	If the field name is an identifier, the current value must be a struct,</span>
<a id="L87"></a><span class="comment">	and there must be a field with that name in the struct. The same lookup</span>
<a id="L88"></a><span class="comment">	rules apply as in the Go language (for instance, the name of an anonymous</span>
<a id="L89"></a><span class="comment">	field is the unqualified type name). The field name denotes the field</span>
<a id="L90"></a><span class="comment">	value in the struct. If the field is not found, formatting is aborted</span>
<a id="L91"></a><span class="comment">	and an error message is returned. (TODO consider changing the semantics</span>
<a id="L92"></a><span class="comment">	such that if a field is not found, it evaluates to nil).</span>

<a id="L94"></a><span class="comment">	The special name &#39;@&#39; denotes the current value.</span>

<a id="L96"></a><span class="comment">	The meaning of the special name &#39;*&#39; depends on the type of the current</span>
<a id="L97"></a><span class="comment">	value:</span>

<a id="L99"></a><span class="comment">		array, slice types   array, slice element (inside {} only, see below)</span>
<a id="L100"></a><span class="comment">		interfaces           value stored in interface</span>
<a id="L101"></a><span class="comment">		pointers             value pointed to by pointer</span>

<a id="L103"></a><span class="comment">	(Implementation restriction: channel, function and map types are not</span>
<a id="L104"></a><span class="comment">	supported due to missing reflection support).</span>

<a id="L106"></a><span class="comment">	Fields are evaluated as follows: If the field value is nil, or an array</span>
<a id="L107"></a><span class="comment">	or slice element does not exist, the result is nil (see below for details</span>
<a id="L108"></a><span class="comment">	on array/slice elements). If the value is not nil the field value is</span>
<a id="L109"></a><span class="comment">	formatted (recursively) using the rule corresponding to its type name,</span>
<a id="L110"></a><span class="comment">	or the alternate rule name, if given.</span>

<a id="L112"></a><span class="comment">	The following example shows a complete format specification for a</span>
<a id="L113"></a><span class="comment">	struct &#39;myPackage.Point&#39;. Assume the package</span>

<a id="L115"></a><span class="comment">		package myPackage  // in directory myDir/myPackage</span>
<a id="L116"></a><span class="comment">		type Point struct {</span>
<a id="L117"></a><span class="comment">			name string;</span>
<a id="L118"></a><span class="comment">			x, y int;</span>
<a id="L119"></a><span class="comment">		}</span>

<a id="L121"></a><span class="comment">	Applying the format specification</span>

<a id="L123"></a><span class="comment">		myPackage &#34;myDir/myPackage&#34;;</span>
<a id="L124"></a><span class="comment">		int = &#34;%d&#34;;</span>
<a id="L125"></a><span class="comment">		hexInt = &#34;0x%x&#34;;</span>
<a id="L126"></a><span class="comment">		string = &#34;---%s---&#34;;</span>
<a id="L127"></a><span class="comment">		myPackage.Point = name &#34;{&#34; x &#34;, &#34; y:hexInt &#34;}&#34;;</span>

<a id="L129"></a><span class="comment">	to the value myPackage.Point{&#34;foo&#34;, 3, 15} results in</span>

<a id="L131"></a><span class="comment">		---foo---{3, 0xf}</span>

<a id="L133"></a><span class="comment">	Finally, an operand may be a grouped, optional, or repeated expression.</span>
<a id="L134"></a><span class="comment">	A grouped expression (&#34;group&#34;) groups a more complex expression (body)</span>
<a id="L135"></a><span class="comment">	so that it can be used in place of a single operand:</span>

<a id="L137"></a><span class="comment">		Group       = &#34;(&#34; [ Indentation &#34;&gt;&gt;&#34; ] Body &#34;)&#34; .</span>
<a id="L138"></a><span class="comment">		Indentation = Expression .</span>
<a id="L139"></a><span class="comment">		Body        = Expression .</span>

<a id="L141"></a><span class="comment">	A group body may be prefixed by an indentation expression followed by &#39;&gt;&gt;&#39;.</span>
<a id="L142"></a><span class="comment">	The indentation expression is applied to the current value like any other</span>
<a id="L143"></a><span class="comment">	expression and the result, if not nil, is appended to the current indentation</span>
<a id="L144"></a><span class="comment">	during the evaluation of the body (see also formatting state, below).</span>

<a id="L146"></a><span class="comment">	An optional expression (&#34;option&#34;) is enclosed in &#39;[]&#39; brackets.</span>

<a id="L148"></a><span class="comment">		Option      = &#34;[&#34; Body &#34;]&#34; .</span>

<a id="L150"></a><span class="comment">	An option evaluates to its body, except that if the body evaluates to nil,</span>
<a id="L151"></a><span class="comment">	the option expression evaluates to an empty []byte. Thus an option&#39;s purpose</span>
<a id="L152"></a><span class="comment">	is to protect the expression containing the option from a nil operand.</span>

<a id="L154"></a><span class="comment">	A repeated expression (&#34;repetition&#34;) is enclosed in &#39;{}&#39; braces.</span>

<a id="L156"></a><span class="comment">		Repetition  = &#34;{&#34; Body [ &#34;/&#34; Separator ] &#34;}&#34; .</span>
<a id="L157"></a><span class="comment">		Separator   = Expression .</span>

<a id="L159"></a><span class="comment">	A repeated expression is evaluated as follows: The body is evaluated</span>
<a id="L160"></a><span class="comment">	repeatedly and its results are concatenated until the body evaluates</span>
<a id="L161"></a><span class="comment">	to nil. The result of the repetition is the (possibly empty) concatenation,</span>
<a id="L162"></a><span class="comment">	but it is never nil. An implicit index is supplied for the evaluation of</span>
<a id="L163"></a><span class="comment">	the body: that index is used to address elements of arrays or slices. If</span>
<a id="L164"></a><span class="comment">	the corresponding elements do not exist, the field denoting the element</span>
<a id="L165"></a><span class="comment">	evaluates to nil (which in turn may terminate the repetition).</span>

<a id="L167"></a><span class="comment">	The body of a repetition may be followed by a &#39;/&#39; and a &#34;separator&#34;</span>
<a id="L168"></a><span class="comment">	expression. If the separator is present, it is invoked between repetitions</span>
<a id="L169"></a><span class="comment">	of the body.</span>

<a id="L171"></a><span class="comment">	The following example shows a complete format specification for formatting</span>
<a id="L172"></a><span class="comment">	a slice of unnamed type. Applying the specification</span>

<a id="L174"></a><span class="comment">		int = &#34;%b&#34;;</span>
<a id="L175"></a><span class="comment">		array = { * / &#34;, &#34; };  // array is the type name for an unnamed slice</span>

<a id="L177"></a><span class="comment">	to the value &#39;[]int{2, 3, 5, 7}&#39; results in</span>

<a id="L179"></a><span class="comment">		10, 11, 101, 111</span>

<a id="L181"></a><span class="comment">	Default rule: If a format rule named &#39;default&#39; is present, it is used for</span>
<a id="L182"></a><span class="comment">	formatting a value if no other rule was found. A common default rule is</span>

<a id="L184"></a><span class="comment">		default = &#34;%v&#34;</span>

<a id="L186"></a><span class="comment">	to provide default formatting for basic types without having to specify</span>
<a id="L187"></a><span class="comment">	a specific rule for each basic type.</span>

<a id="L189"></a><span class="comment">	Global separator rule: If a format rule named &#39;/&#39; is present, it is</span>
<a id="L190"></a><span class="comment">	invoked with the current value between literals. If the separator</span>
<a id="L191"></a><span class="comment">	expression evaluates to nil, it is ignored.</span>

<a id="L193"></a><span class="comment">	For instance, a global separator rule may be used to punctuate a sequence</span>
<a id="L194"></a><span class="comment">	of values with commas. The rules:</span>

<a id="L196"></a><span class="comment">		default = &#34;%v&#34;;</span>
<a id="L197"></a><span class="comment">		/ = &#34;, &#34;;</span>

<a id="L199"></a><span class="comment">	will format an argument list by printing each one in its default format,</span>
<a id="L200"></a><span class="comment">	separated by a comma and a space.</span>
<a id="L201"></a><span class="comment">*/</span>
<a id="L202"></a>package datafmt

<a id="L204"></a>import (
    <a id="L205"></a>&#34;bytes&#34;;
    <a id="L206"></a>&#34;fmt&#34;;
    <a id="L207"></a>&#34;go/token&#34;;
    <a id="L208"></a>&#34;io&#34;;
    <a id="L209"></a>&#34;os&#34;;
    <a id="L210"></a>&#34;reflect&#34;;
    <a id="L211"></a>&#34;runtime&#34;;
<a id="L212"></a>)


<a id="L215"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L216"></a><span class="comment">// Format representation</span>

<a id="L218"></a><span class="comment">// Custom formatters implement the Formatter function type.</span>
<a id="L219"></a><span class="comment">// A formatter is invoked with the current formatting state, the</span>
<a id="L220"></a><span class="comment">// value to format, and the rule name under which the formatter</span>
<a id="L221"></a><span class="comment">// was installed (the same formatter function may be installed</span>
<a id="L222"></a><span class="comment">// under different names). The formatter may access the current state</span>
<a id="L223"></a><span class="comment">// to guide formatting and use State.Write to append to the state&#39;s</span>
<a id="L224"></a><span class="comment">// output.</span>
<a id="L225"></a><span class="comment">//</span>
<a id="L226"></a><span class="comment">// A formatter must return a boolean value indicating if it evaluated</span>
<a id="L227"></a><span class="comment">// to a non-nil value (true), or a nil value (false).</span>
<a id="L228"></a><span class="comment">//</span>
<a id="L229"></a>type Formatter func(state *State, value interface{}, ruleName string) bool


<a id="L232"></a><span class="comment">// A FormatterMap is a set of custom formatters.</span>
<a id="L233"></a><span class="comment">// It maps a rule name to a formatter function.</span>
<a id="L234"></a><span class="comment">//</span>
<a id="L235"></a>type FormatterMap map[string]Formatter


<a id="L238"></a><span class="comment">// A parsed format expression is built from the following nodes.</span>
<a id="L239"></a><span class="comment">//</span>
<a id="L240"></a>type (
    <a id="L241"></a>expr interface{};

    <a id="L243"></a>alternatives []expr; <span class="comment">// x | y | z</span>

    <a id="L245"></a>sequence []expr; <span class="comment">// x y z</span>

    <a id="L247"></a>literal [][]byte; <span class="comment">// a list of string segments, possibly starting with &#39;%&#39;</span>

    <a id="L249"></a>field struct {
        <a id="L250"></a>fieldName string; <span class="comment">// including &#34;@&#34;, &#34;*&#34;</span>
        <a id="L251"></a>ruleName  string; <span class="comment">// &#34;&#34; if no rule name specified</span>
    <a id="L252"></a>};

    <a id="L254"></a>group struct {
        <a id="L255"></a>indent, body expr; <span class="comment">// (indent &gt;&gt; body)</span>
    <a id="L256"></a>};

    <a id="L258"></a>option struct {
        <a id="L259"></a>body expr; <span class="comment">// [body]</span>
    <a id="L260"></a>};

    <a id="L262"></a>repetition struct {
        <a id="L263"></a>body, separator expr; <span class="comment">// {body / separator}</span>
    <a id="L264"></a>};

    <a id="L266"></a>custom struct {
        <a id="L267"></a>ruleName string;
        <a id="L268"></a>fun      Formatter;
    <a id="L269"></a>};
<a id="L270"></a>)


<a id="L273"></a><span class="comment">// A Format is the result of parsing a format specification.</span>
<a id="L274"></a><span class="comment">// The format may be applied repeatedly to format values.</span>
<a id="L275"></a><span class="comment">//</span>
<a id="L276"></a>type Format map[string]expr


<a id="L279"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L280"></a><span class="comment">// Formatting</span>

<a id="L282"></a><span class="comment">// An application-specific environment may be provided to Format.Apply;</span>
<a id="L283"></a><span class="comment">// the environment is available inside custom formatters via State.Env().</span>
<a id="L284"></a><span class="comment">// Environments must implement copying; the Copy method must return an</span>
<a id="L285"></a><span class="comment">// complete copy of the receiver. This is necessary so that the formatter</span>
<a id="L286"></a><span class="comment">// can save and restore an environment (in case of an absent expression).</span>
<a id="L287"></a><span class="comment">//</span>
<a id="L288"></a><span class="comment">// If the Environment doesn&#39;t change during formatting (this is under</span>
<a id="L289"></a><span class="comment">// control of the custom formatters), the Copy function can simply return</span>
<a id="L290"></a><span class="comment">// the receiver, and thus can be very light-weight.</span>
<a id="L291"></a><span class="comment">//</span>
<a id="L292"></a>type Environment interface {
    <a id="L293"></a>Copy() Environment;
<a id="L294"></a>}


<a id="L297"></a><span class="comment">// State represents the current formatting state.</span>
<a id="L298"></a><span class="comment">// It is provided as argument to custom formatters.</span>
<a id="L299"></a><span class="comment">//</span>
<a id="L300"></a>type State struct {
    <a id="L301"></a>fmt       Format;         <span class="comment">// format in use</span>
    <a id="L302"></a>env       Environment;    <span class="comment">// user-supplied environment</span>
    <a id="L303"></a>errors    chan os.Error;  <span class="comment">// not chan *Error (errors &lt;- nil would be wrong!)</span>
    <a id="L304"></a>hasOutput bool;           <span class="comment">// true after the first literal has been written</span>
    <a id="L305"></a>indent    bytes.Buffer;   <span class="comment">// current indentation</span>
    <a id="L306"></a>output    bytes.Buffer;   <span class="comment">// format output</span>
    <a id="L307"></a>linePos   token.Position; <span class="comment">// position of line beginning (Column == 0)</span>
    <a id="L308"></a>default_  expr;           <span class="comment">// possibly nil</span>
    <a id="L309"></a>separator expr;           <span class="comment">// possibly nil</span>
<a id="L310"></a>}


<a id="L313"></a>func newState(fmt Format, env Environment, errors chan os.Error) *State {
    <a id="L314"></a>s := new(State);
    <a id="L315"></a>s.fmt = fmt;
    <a id="L316"></a>s.env = env;
    <a id="L317"></a>s.errors = errors;
    <a id="L318"></a>s.linePos = token.Position{Line: 1};

    <a id="L320"></a><span class="comment">// if we have a default rule, cache it&#39;s expression for fast access</span>
    <a id="L321"></a>if x, found := fmt[&#34;default&#34;]; found {
        <a id="L322"></a>s.default_ = x
    <a id="L323"></a>}

    <a id="L325"></a><span class="comment">// if we have a global separator rule, cache it&#39;s expression for fast access</span>
    <a id="L326"></a>if x, found := fmt[&#34;/&#34;]; found {
        <a id="L327"></a>s.separator = x
    <a id="L328"></a>}

    <a id="L330"></a>return s;
<a id="L331"></a>}


<a id="L334"></a><span class="comment">// Env returns the environment passed to Format.Apply.</span>
<a id="L335"></a>func (s *State) Env() interface{} { return s.env }


<a id="L338"></a><span class="comment">// LinePos returns the position of the current line beginning</span>
<a id="L339"></a><span class="comment">// in the state&#39;s output buffer. Line numbers start at 1.</span>
<a id="L340"></a><span class="comment">//</span>
<a id="L341"></a>func (s *State) LinePos() token.Position { return s.linePos }


<a id="L344"></a><span class="comment">// Pos returns the position of the next byte to be written to the</span>
<a id="L345"></a><span class="comment">// output buffer. Line numbers start at 1.</span>
<a id="L346"></a><span class="comment">//</span>
<a id="L347"></a>func (s *State) Pos() token.Position {
    <a id="L348"></a>offs := s.output.Len();
    <a id="L349"></a>return token.Position{Line: s.linePos.Line, Column: offs - s.linePos.Offset, Offset: offs};
<a id="L350"></a>}


<a id="L353"></a><span class="comment">// Write writes data to the output buffer, inserting the indentation</span>
<a id="L354"></a><span class="comment">// string after each newline or form feed character. It cannot return an error.</span>
<a id="L355"></a><span class="comment">//</span>
<a id="L356"></a>func (s *State) Write(data []byte) (int, os.Error) {
    <a id="L357"></a>n := 0;
    <a id="L358"></a>i0 := 0;
    <a id="L359"></a>for i, ch := range data {
        <a id="L360"></a>if ch == &#39;\n&#39; || ch == &#39;\f&#39; {
            <a id="L361"></a><span class="comment">// write text segment and indentation</span>
            <a id="L362"></a>n1, _ := s.output.Write(data[i0 : i+1]);
            <a id="L363"></a>n2, _ := s.output.Write(s.indent.Bytes());
            <a id="L364"></a>n += n1 + n2;
            <a id="L365"></a>i0 = i + 1;
            <a id="L366"></a>s.linePos.Offset = s.output.Len();
            <a id="L367"></a>s.linePos.Line++;
        <a id="L368"></a>}
    <a id="L369"></a>}
    <a id="L370"></a>n3, _ := s.output.Write(data[i0:len(data)]);
    <a id="L371"></a>return n + n3, nil;
<a id="L372"></a>}


<a id="L375"></a>type checkpoint struct {
    <a id="L376"></a>env       Environment;
    <a id="L377"></a>hasOutput bool;
    <a id="L378"></a>outputLen int;
    <a id="L379"></a>linePos   token.Position;
<a id="L380"></a>}


<a id="L383"></a>func (s *State) save() checkpoint {
    <a id="L384"></a>saved := checkpoint{nil, s.hasOutput, s.output.Len(), s.linePos};
    <a id="L385"></a>if s.env != nil {
        <a id="L386"></a>saved.env = s.env.Copy()
    <a id="L387"></a>}
    <a id="L388"></a>return saved;
<a id="L389"></a>}


<a id="L392"></a>func (s *State) restore(m checkpoint) {
    <a id="L393"></a>s.env = m.env;
    <a id="L394"></a>s.output.Truncate(m.outputLen);
<a id="L395"></a>}


<a id="L398"></a>func (s *State) error(msg string) {
    <a id="L399"></a>s.errors &lt;- os.NewError(msg);
    <a id="L400"></a>runtime.Goexit();
<a id="L401"></a>}


<a id="L404"></a><span class="comment">// TODO At the moment, unnamed types are simply mapped to the default</span>
<a id="L405"></a><span class="comment">//      names below. For instance, all unnamed arrays are mapped to</span>
<a id="L406"></a><span class="comment">//      &#39;array&#39; which is not really sufficient. Eventually one may want</span>
<a id="L407"></a><span class="comment">//      to be able to specify rules for say an unnamed slice of T.</span>
<a id="L408"></a><span class="comment">//</span>

<a id="L410"></a>func typename(typ reflect.Type) string {
    <a id="L411"></a>switch typ.(type) {
    <a id="L412"></a>case *reflect.ArrayType:
        <a id="L413"></a>return &#34;array&#34;
    <a id="L414"></a>case *reflect.SliceType:
        <a id="L415"></a>return &#34;array&#34;
    <a id="L416"></a>case *reflect.ChanType:
        <a id="L417"></a>return &#34;chan&#34;
    <a id="L418"></a>case *reflect.DotDotDotType:
        <a id="L419"></a>return &#34;ellipsis&#34;
    <a id="L420"></a>case *reflect.FuncType:
        <a id="L421"></a>return &#34;func&#34;
    <a id="L422"></a>case *reflect.InterfaceType:
        <a id="L423"></a>return &#34;interface&#34;
    <a id="L424"></a>case *reflect.MapType:
        <a id="L425"></a>return &#34;map&#34;
    <a id="L426"></a>case *reflect.PtrType:
        <a id="L427"></a>return &#34;ptr&#34;
    <a id="L428"></a>}
    <a id="L429"></a>return typ.String();
<a id="L430"></a>}

<a id="L432"></a>func (s *State) getFormat(name string) expr {
    <a id="L433"></a>if fexpr, found := s.fmt[name]; found {
        <a id="L434"></a>return fexpr
    <a id="L435"></a>}

    <a id="L437"></a>if s.default_ != nil {
        <a id="L438"></a>return s.default_
    <a id="L439"></a>}

    <a id="L441"></a>s.error(fmt.Sprintf(&#34;no format rule for type: &#39;%s&#39;&#34;, name));
    <a id="L442"></a>return nil;
<a id="L443"></a>}


<a id="L446"></a><span class="comment">// eval applies a format expression fexpr to a value. If the expression</span>
<a id="L447"></a><span class="comment">// evaluates internally to a non-nil []byte, that slice is appended to</span>
<a id="L448"></a><span class="comment">// the state&#39;s output buffer and eval returns true. Otherwise, eval</span>
<a id="L449"></a><span class="comment">// returns false and the state remains unchanged.</span>
<a id="L450"></a><span class="comment">//</span>
<a id="L451"></a>func (s *State) eval(fexpr expr, value reflect.Value, index int) bool {
    <a id="L452"></a><span class="comment">// an empty format expression always evaluates</span>
    <a id="L453"></a><span class="comment">// to a non-nil (but empty) []byte</span>
    <a id="L454"></a>if fexpr == nil {
        <a id="L455"></a>return true
    <a id="L456"></a>}

    <a id="L458"></a>switch t := fexpr.(type) {
    <a id="L459"></a>case alternatives:
        <a id="L460"></a><span class="comment">// append the result of the first alternative that evaluates to</span>
        <a id="L461"></a><span class="comment">// a non-nil []byte to the state&#39;s output</span>
        <a id="L462"></a>mark := s.save();
        <a id="L463"></a>for _, x := range t {
            <a id="L464"></a>if s.eval(x, value, index) {
                <a id="L465"></a>return true
            <a id="L466"></a>}
            <a id="L467"></a>s.restore(mark);
        <a id="L468"></a>}
        <a id="L469"></a>return false;

    <a id="L471"></a>case sequence:
        <a id="L472"></a><span class="comment">// append the result of all operands to the state&#39;s output</span>
        <a id="L473"></a><span class="comment">// unless a nil result is encountered</span>
        <a id="L474"></a>mark := s.save();
        <a id="L475"></a>for _, x := range t {
            <a id="L476"></a>if !s.eval(x, value, index) {
                <a id="L477"></a>s.restore(mark);
                <a id="L478"></a>return false;
            <a id="L479"></a>}
        <a id="L480"></a>}
        <a id="L481"></a>return true;

    <a id="L483"></a>case literal:
        <a id="L484"></a><span class="comment">// write separator, if any</span>
        <a id="L485"></a>if s.hasOutput {
            <a id="L486"></a><span class="comment">// not the first literal</span>
            <a id="L487"></a>if s.separator != nil {
                <a id="L488"></a>sep := s.separator; <span class="comment">// save current separator</span>
                <a id="L489"></a>s.separator = nil;  <span class="comment">// and disable it (avoid recursion)</span>
                <a id="L490"></a>mark := s.save();
                <a id="L491"></a>if !s.eval(sep, value, index) {
                    <a id="L492"></a>s.restore(mark)
                <a id="L493"></a>}
                <a id="L494"></a>s.separator = sep; <span class="comment">// enable it again</span>
            <a id="L495"></a>}
        <a id="L496"></a>}
        <a id="L497"></a>s.hasOutput = true;
        <a id="L498"></a><span class="comment">// write literal segments</span>
        <a id="L499"></a>for _, lit := range t {
            <a id="L500"></a>if len(lit) &gt; 1 &amp;&amp; lit[0] == &#39;%&#39; {
                <a id="L501"></a><span class="comment">// segment contains a %-format at the beginning</span>
                <a id="L502"></a>if lit[1] == &#39;%&#39; {
                    <a id="L503"></a><span class="comment">// &#34;%%&#34; is printed as a single &#34;%&#34;</span>
                    <a id="L504"></a>s.Write(lit[1:len(lit)])
                <a id="L505"></a>} else {
                    <a id="L506"></a><span class="comment">// use s instead of s.output to get indentation right</span>
                    <a id="L507"></a>fmt.Fprintf(s, string(lit), value.Interface())
                <a id="L508"></a>}
            <a id="L509"></a>} else {
                <a id="L510"></a><span class="comment">// segment contains no %-formats</span>
                <a id="L511"></a>s.Write(lit)
            <a id="L512"></a>}
        <a id="L513"></a>}
        <a id="L514"></a>return true; <span class="comment">// a literal never evaluates to nil</span>

    <a id="L516"></a>case *field:
        <a id="L517"></a><span class="comment">// determine field value</span>
        <a id="L518"></a>switch t.fieldName {
        <a id="L519"></a>case &#34;@&#34;:
            <a id="L520"></a><span class="comment">// field value is current value</span>

        <a id="L522"></a>case &#34;*&#34;:
            <a id="L523"></a><span class="comment">// indirection: operation is type-specific</span>
            <a id="L524"></a>switch v := value.(type) {
            <a id="L525"></a>case *reflect.ArrayValue:
                <a id="L526"></a>if v.Len() &lt;= index {
                    <a id="L527"></a>return false
                <a id="L528"></a>}
                <a id="L529"></a>value = v.Elem(index);

            <a id="L531"></a>case *reflect.SliceValue:
                <a id="L532"></a>if v.IsNil() || v.Len() &lt;= index {
                    <a id="L533"></a>return false
                <a id="L534"></a>}
                <a id="L535"></a>value = v.Elem(index);

            <a id="L537"></a>case *reflect.MapValue:
                <a id="L538"></a>s.error(&#34;reflection support for maps incomplete&#34;)

            <a id="L540"></a>case *reflect.PtrValue:
                <a id="L541"></a>if v.IsNil() {
                    <a id="L542"></a>return false
                <a id="L543"></a>}
                <a id="L544"></a>value = v.Elem();

            <a id="L546"></a>case *reflect.InterfaceValue:
                <a id="L547"></a>if v.IsNil() {
                    <a id="L548"></a>return false
                <a id="L549"></a>}
                <a id="L550"></a>value = v.Elem();

            <a id="L552"></a>case *reflect.ChanValue:
                <a id="L553"></a>s.error(&#34;reflection support for chans incomplete&#34;)

            <a id="L555"></a>case *reflect.FuncValue:
                <a id="L556"></a>s.error(&#34;reflection support for funcs incomplete&#34;)

            <a id="L558"></a>default:
                <a id="L559"></a>s.error(fmt.Sprintf(&#34;error: * does not apply to `%s`&#34;, value.Type()))
            <a id="L560"></a>}

        <a id="L562"></a>default:
            <a id="L563"></a><span class="comment">// value is value of named field</span>
            <a id="L564"></a>var field reflect.Value;
            <a id="L565"></a>if sval, ok := value.(*reflect.StructValue); ok {
                <a id="L566"></a>field = sval.FieldByName(t.fieldName);
                <a id="L567"></a>if field == nil {
                    <a id="L568"></a><span class="comment">// TODO consider just returning false in this case</span>
                    <a id="L569"></a>s.error(fmt.Sprintf(&#34;error: no field `%s` in `%s`&#34;, t.fieldName, value.Type()))
                <a id="L570"></a>}
            <a id="L571"></a>}
            <a id="L572"></a>value = field;
        <a id="L573"></a>}

        <a id="L575"></a><span class="comment">// determine rule</span>
        <a id="L576"></a>ruleName := t.ruleName;
        <a id="L577"></a>if ruleName == &#34;&#34; {
            <a id="L578"></a><span class="comment">// no alternate rule name, value type determines rule</span>
            <a id="L579"></a>ruleName = typename(value.Type())
        <a id="L580"></a>}
        <a id="L581"></a>fexpr = s.getFormat(ruleName);

        <a id="L583"></a>mark := s.save();
        <a id="L584"></a>if !s.eval(fexpr, value, index) {
            <a id="L585"></a>s.restore(mark);
            <a id="L586"></a>return false;
        <a id="L587"></a>}
        <a id="L588"></a>return true;

    <a id="L590"></a>case *group:
        <a id="L591"></a><span class="comment">// remember current indentation</span>
        <a id="L592"></a>indentLen := s.indent.Len();

        <a id="L594"></a><span class="comment">// update current indentation</span>
        <a id="L595"></a>mark := s.save();
        <a id="L596"></a>s.eval(t.indent, value, index);
        <a id="L597"></a><span class="comment">// if the indentation evaluates to nil, the state&#39;s output buffer</span>
        <a id="L598"></a><span class="comment">// didn&#39;t change - either way it&#39;s ok to append the difference to</span>
        <a id="L599"></a><span class="comment">// the current identation</span>
        <a id="L600"></a>s.indent.Write(s.output.Bytes()[mark.outputLen:s.output.Len()]);
        <a id="L601"></a>s.restore(mark);

        <a id="L603"></a><span class="comment">// format group body</span>
        <a id="L604"></a>mark = s.save();
        <a id="L605"></a>b := true;
        <a id="L606"></a>if !s.eval(t.body, value, index) {
            <a id="L607"></a>s.restore(mark);
            <a id="L608"></a>b = false;
        <a id="L609"></a>}

        <a id="L611"></a><span class="comment">// reset indentation</span>
        <a id="L612"></a>s.indent.Truncate(indentLen);
        <a id="L613"></a>return b;

    <a id="L615"></a>case *option:
        <a id="L616"></a><span class="comment">// evaluate the body and append the result to the state&#39;s output</span>
        <a id="L617"></a><span class="comment">// buffer unless the result is nil</span>
        <a id="L618"></a>mark := s.save();
        <a id="L619"></a>if !s.eval(t.body, value, 0) { <span class="comment">// TODO is 0 index correct?</span>
            <a id="L620"></a>s.restore(mark)
        <a id="L621"></a>}
        <a id="L622"></a>return true; <span class="comment">// an option never evaluates to nil</span>

    <a id="L624"></a>case *repetition:
        <a id="L625"></a><span class="comment">// evaluate the body and append the result to the state&#39;s output</span>
        <a id="L626"></a><span class="comment">// buffer until a result is nil</span>
        <a id="L627"></a>for i := 0; ; i++ {
            <a id="L628"></a>mark := s.save();
            <a id="L629"></a><span class="comment">// write separator, if any</span>
            <a id="L630"></a>if i &gt; 0 &amp;&amp; t.separator != nil {
                <a id="L631"></a><span class="comment">// nil result from separator is ignored</span>
                <a id="L632"></a>mark := s.save();
                <a id="L633"></a>if !s.eval(t.separator, value, i) {
                    <a id="L634"></a>s.restore(mark)
                <a id="L635"></a>}
            <a id="L636"></a>}
            <a id="L637"></a>if !s.eval(t.body, value, i) {
                <a id="L638"></a>s.restore(mark);
                <a id="L639"></a>break;
            <a id="L640"></a>}
        <a id="L641"></a>}
        <a id="L642"></a>return true; <span class="comment">// a repetition never evaluates to nil</span>

    <a id="L644"></a>case *custom:
        <a id="L645"></a><span class="comment">// invoke the custom formatter to obtain the result</span>
        <a id="L646"></a>mark := s.save();
        <a id="L647"></a>if !t.fun(s, value.Interface(), t.ruleName) {
            <a id="L648"></a>s.restore(mark);
            <a id="L649"></a>return false;
        <a id="L650"></a>}
        <a id="L651"></a>return true;
    <a id="L652"></a>}

    <a id="L654"></a>panic(&#34;unreachable&#34;);
    <a id="L655"></a>return false;
<a id="L656"></a>}


<a id="L659"></a><span class="comment">// Eval formats each argument according to the format</span>
<a id="L660"></a><span class="comment">// f and returns the resulting []byte and os.Error. If</span>
<a id="L661"></a><span class="comment">// an error occured, the []byte contains the partially</span>
<a id="L662"></a><span class="comment">// formatted result. An environment env may be passed</span>
<a id="L663"></a><span class="comment">// in which is available in custom formatters through</span>
<a id="L664"></a><span class="comment">// the state parameter.</span>
<a id="L665"></a><span class="comment">//</span>
<a id="L666"></a>func (f Format) Eval(env Environment, args ...) ([]byte, os.Error) {
    <a id="L667"></a>if f == nil {
        <a id="L668"></a>return nil, os.NewError(&#34;format is nil&#34;)
    <a id="L669"></a>}

    <a id="L671"></a>errors := make(chan os.Error);
    <a id="L672"></a>s := newState(f, env, errors);

    <a id="L674"></a>go func() {
        <a id="L675"></a>value := reflect.NewValue(args).(*reflect.StructValue);
        <a id="L676"></a>for i := 0; i &lt; value.NumField(); i++ {
            <a id="L677"></a>fld := value.Field(i);
            <a id="L678"></a>mark := s.save();
            <a id="L679"></a>if !s.eval(s.getFormat(typename(fld.Type())), fld, 0) { <span class="comment">// TODO is 0 index correct?</span>
                <a id="L680"></a>s.restore(mark)
            <a id="L681"></a>}
        <a id="L682"></a>}
        <a id="L683"></a>errors &lt;- nil; <span class="comment">// no errors</span>
    <a id="L684"></a>}();

    <a id="L686"></a>err := &lt;-errors;
    <a id="L687"></a>return s.output.Bytes(), err;
<a id="L688"></a>}


<a id="L691"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L692"></a><span class="comment">// Convenience functions</span>

<a id="L694"></a><span class="comment">// Fprint formats each argument according to the format f</span>
<a id="L695"></a><span class="comment">// and writes to w. The result is the total number of bytes</span>
<a id="L696"></a><span class="comment">// written and an os.Error, if any.</span>
<a id="L697"></a><span class="comment">//</span>
<a id="L698"></a>func (f Format) Fprint(w io.Writer, env Environment, args ...) (int, os.Error) {
    <a id="L699"></a>data, err := f.Eval(env, args);
    <a id="L700"></a>if err != nil {
        <a id="L701"></a><span class="comment">// TODO should we print partial result in case of error?</span>
        <a id="L702"></a>return 0, err
    <a id="L703"></a>}
    <a id="L704"></a>return w.Write(data);
<a id="L705"></a>}


<a id="L708"></a><span class="comment">// Print formats each argument according to the format f</span>
<a id="L709"></a><span class="comment">// and writes to standard output. The result is the total</span>
<a id="L710"></a><span class="comment">// number of bytes written and an os.Error, if any.</span>
<a id="L711"></a><span class="comment">//</span>
<a id="L712"></a>func (f Format) Print(args ...) (int, os.Error) {
    <a id="L713"></a>return f.Fprint(os.Stdout, nil, args)
<a id="L714"></a>}


<a id="L717"></a><span class="comment">// Sprint formats each argument according to the format f</span>
<a id="L718"></a><span class="comment">// and returns the resulting string. If an error occurs</span>
<a id="L719"></a><span class="comment">// during formatting, the result string contains the</span>
<a id="L720"></a><span class="comment">// partially formatted result followed by an error message.</span>
<a id="L721"></a><span class="comment">//</span>
<a id="L722"></a>func (f Format) Sprint(args ...) string {
    <a id="L723"></a>var buf bytes.Buffer;
    <a id="L724"></a>_, err := f.Fprint(&amp;buf, nil, args);
    <a id="L725"></a>if err != nil {
        <a id="L726"></a>fmt.Fprintf(&amp;buf, &#34;--- Sprint(%s) failed: %v&#34;, fmt.Sprint(args), err)
    <a id="L727"></a>}
    <a id="L728"></a>return buf.String();
<a id="L729"></a>}
</pre>

</div>

<div id="footer">
<p>Except as noted, this content is
   licensed under <a href="http://creativecommons.org/licenses/by/3.0/">
   Creative Commons Attribution 3.0</a>.
</div>

<script type="text/javascript">
var gaJsHost = (("https:" == document.location.protocol) ? "https://ssl." : "http://www.");
document.write(unescape("%3Cscript src='" + gaJsHost + "google-analytics.com/ga.js' type='text/javascript'%3E%3C/script%3E"));
</script>
<script type="text/javascript">
var pageTracker = _gat._getTracker("UA-11222381-2");
pageTracker._trackPageview();
</script>
</body>
</html>
<!-- generated at Thu Nov 12 15:42:51 PST 2009 -->
