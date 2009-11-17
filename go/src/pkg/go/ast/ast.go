<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/go/ast/ast.go</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/go/ast/ast.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The AST package declares the types used to represent</span>
<a id="L6"></a><span class="comment">// syntax trees for Go packages.</span>
<a id="L7"></a><span class="comment">//</span>
<a id="L8"></a>package ast

<a id="L10"></a>import (
    <a id="L11"></a>&#34;go/token&#34;;
    <a id="L12"></a>&#34;unicode&#34;;
    <a id="L13"></a>&#34;utf8&#34;;
<a id="L14"></a>)


<a id="L17"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L18"></a><span class="comment">// Interfaces</span>
<a id="L19"></a><span class="comment">//</span>
<a id="L20"></a><span class="comment">// There are 3 main classes of nodes: Expressions and type nodes,</span>
<a id="L21"></a><span class="comment">// statement nodes, and declaration nodes. The node names usually</span>
<a id="L22"></a><span class="comment">// match the corresponding Go spec production names to which they</span>
<a id="L23"></a><span class="comment">// correspond. The node fields correspond to the individual parts</span>
<a id="L24"></a><span class="comment">// of the respective productions.</span>
<a id="L25"></a><span class="comment">//</span>
<a id="L26"></a><span class="comment">// All nodes contain position information marking the beginning of</span>
<a id="L27"></a><span class="comment">// the corresponding source text segment; it is accessible via the</span>
<a id="L28"></a><span class="comment">// Pos accessor method. Nodes may contain additional position info</span>
<a id="L29"></a><span class="comment">// for language constructs where comments may be found between parts</span>
<a id="L30"></a><span class="comment">// of the construct (typically any larger, parenthesized subpart).</span>
<a id="L31"></a><span class="comment">// That position information is needed to properly position comments</span>
<a id="L32"></a><span class="comment">// when printing the construct.</span>


<a id="L35"></a><span class="comment">// All node types implement the Node interface.</span>
<a id="L36"></a>type Node interface {
    <a id="L37"></a><span class="comment">// Pos returns the (beginning) position of the node.</span>
    <a id="L38"></a>Pos() token.Position;
<a id="L39"></a>}


<a id="L42"></a><span class="comment">// All expression nodes implement the Expr interface.</span>
<a id="L43"></a>type Expr interface {
    <a id="L44"></a>Node;
    <a id="L45"></a>exprNode();
<a id="L46"></a>}


<a id="L49"></a><span class="comment">// All statement nodes implement the Stmt interface.</span>
<a id="L50"></a>type Stmt interface {
    <a id="L51"></a>Node;
    <a id="L52"></a>stmtNode();
<a id="L53"></a>}


<a id="L56"></a><span class="comment">// All declaration nodes implement the Decl interface.</span>
<a id="L57"></a>type Decl interface {
    <a id="L58"></a>Node;
    <a id="L59"></a>declNode();
<a id="L60"></a>}


<a id="L63"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L64"></a><span class="comment">// Comments</span>

<a id="L66"></a><span class="comment">// A Comment node represents a single //-style or /*-style comment.</span>
<a id="L67"></a>type Comment struct {
    <a id="L68"></a>token.Position;         <span class="comment">// beginning position of the comment</span>
    <a id="L69"></a>Text            []byte; <span class="comment">// comment text (excluding &#39;\n&#39; for //-style comments)</span>
<a id="L70"></a>}


<a id="L73"></a><span class="comment">// A CommentGroup represents a sequence of comments</span>
<a id="L74"></a><span class="comment">// with no other tokens and no empty lines between.</span>
<a id="L75"></a><span class="comment">//</span>
<a id="L76"></a>type CommentGroup struct {
    <a id="L77"></a>List []*Comment;
    <a id="L78"></a>Next *CommentGroup; <span class="comment">// next comment group in source order</span>
<a id="L79"></a>}


<a id="L82"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L83"></a><span class="comment">// Expressions and types</span>

<a id="L85"></a><span class="comment">// A Field represents a Field declaration list in a struct type,</span>
<a id="L86"></a><span class="comment">// a method list in an interface type, or a parameter/result declaration</span>
<a id="L87"></a><span class="comment">// in a signature.</span>
<a id="L88"></a><span class="comment">//</span>
<a id="L89"></a>type Field struct {
    <a id="L90"></a>Doc     *CommentGroup; <span class="comment">// associated documentation; or nil</span>
    <a id="L91"></a>Names   []*Ident;      <span class="comment">// field/method/parameter names; or nil if anonymous field</span>
    <a id="L92"></a>Type    Expr;          <span class="comment">// field/method/parameter type</span>
    <a id="L93"></a>Tag     []*BasicLit;   <span class="comment">// field tag; or nil</span>
    <a id="L94"></a>Comment *CommentGroup; <span class="comment">// line comments; or nil</span>
<a id="L95"></a>}


<a id="L98"></a>func (f *Field) Pos() token.Position {
    <a id="L99"></a>if len(f.Names) &gt; 0 {
        <a id="L100"></a>return f.Names[0].Pos()
    <a id="L101"></a>}
    <a id="L102"></a>return f.Type.Pos();
<a id="L103"></a>}


<a id="L106"></a><span class="comment">// An expression is represented by a tree consisting of one</span>
<a id="L107"></a><span class="comment">// or more of the following concrete expression nodes.</span>
<a id="L108"></a><span class="comment">//</span>
<a id="L109"></a>type (
    <a id="L110"></a><span class="comment">// A BadExpr node is a placeholder for expressions containing</span>
    <a id="L111"></a><span class="comment">// syntax errors for which no correct expression nodes can be</span>
    <a id="L112"></a><span class="comment">// created.</span>
    <a id="L113"></a><span class="comment">//</span>
    <a id="L114"></a>BadExpr struct {
        <a id="L115"></a>token.Position; <span class="comment">// beginning position of bad expression</span>
    <a id="L116"></a>};

    <a id="L118"></a><span class="comment">// An Ident node represents an identifier.</span>
    <a id="L119"></a>Ident struct {
        <a id="L120"></a>token.Position;         <span class="comment">// identifier position</span>
        <a id="L121"></a>Value           string; <span class="comment">// identifier string (e.g. foobar)</span>
    <a id="L122"></a>};

    <a id="L124"></a><span class="comment">// An Ellipsis node stands for the &#34;...&#34; type in a</span>
    <a id="L125"></a><span class="comment">// parameter list or the &#34;...&#34; length in an array type.</span>
    <a id="L126"></a><span class="comment">//</span>
    <a id="L127"></a>Ellipsis struct {
        <a id="L128"></a>token.Position; <span class="comment">// position of &#34;...&#34;</span>
    <a id="L129"></a>};

    <a id="L131"></a><span class="comment">// A BasicLit node represents a literal of basic type.</span>
    <a id="L132"></a>BasicLit struct {
        <a id="L133"></a>token.Position;              <span class="comment">// literal position</span>
        <a id="L134"></a>Kind            token.Token; <span class="comment">//  token.INT, token.FLOAT, token.CHAR, or token.STRING</span>
        <a id="L135"></a>Value           []byte;      <span class="comment">// literal string; e.g. 42, 0x7f, 3.14, 1e-9, &#39;a&#39;, &#39;\x7f&#39;, &#34;foo&#34; or `\m\n\o`</span>
    <a id="L136"></a>};

    <a id="L138"></a><span class="comment">// A StringList node represents a sequence of adjacent string literals.</span>
    <a id="L139"></a><span class="comment">// A single string literal (common case) is represented by a BasicLit</span>
    <a id="L140"></a><span class="comment">// node; StringList nodes are used only if there are two or more string</span>
    <a id="L141"></a><span class="comment">// literals in a sequence.</span>
    <a id="L142"></a><span class="comment">//</span>
    <a id="L143"></a>StringList struct {
        <a id="L144"></a>Strings []*BasicLit; <span class="comment">// list of strings, len(Strings) &gt; 1</span>
    <a id="L145"></a>};

    <a id="L147"></a><span class="comment">// A FuncLit node represents a function literal.</span>
    <a id="L148"></a>FuncLit struct {
        <a id="L149"></a>Type *FuncType;  <span class="comment">// function type</span>
        <a id="L150"></a>Body *BlockStmt; <span class="comment">// function body</span>
    <a id="L151"></a>};

    <a id="L153"></a><span class="comment">// A CompositeLit node represents a composite literal.</span>
    <a id="L154"></a><span class="comment">//</span>
    <a id="L155"></a>CompositeLit struct {
        <a id="L156"></a>Type   Expr;           <span class="comment">// literal type</span>
        <a id="L157"></a>Lbrace token.Position; <span class="comment">// position of &#34;{&#34;</span>
        <a id="L158"></a>Elts   []Expr;         <span class="comment">// list of composite elements</span>
        <a id="L159"></a>Rbrace token.Position; <span class="comment">// position of &#34;}&#34;</span>
    <a id="L160"></a>};

    <a id="L162"></a><span class="comment">// A ParenExpr node represents a parenthesized expression.</span>
    <a id="L163"></a>ParenExpr struct {
        <a id="L164"></a>token.Position;                 <span class="comment">// position of &#34;(&#34;</span>
        <a id="L165"></a>X               Expr;           <span class="comment">// parenthesized expression</span>
        <a id="L166"></a>Rparen          token.Position; <span class="comment">// position of &#34;)&#34;</span>
    <a id="L167"></a>};

    <a id="L169"></a><span class="comment">// A SelectorExpr node represents an expression followed by a selector.</span>
    <a id="L170"></a>SelectorExpr struct {
        <a id="L171"></a>X   Expr;   <span class="comment">// expression</span>
        <a id="L172"></a>Sel *Ident; <span class="comment">// field selector</span>
    <a id="L173"></a>};

    <a id="L175"></a><span class="comment">// An IndexExpr node represents an expression followed by an index or slice.</span>
    <a id="L176"></a>IndexExpr struct {
        <a id="L177"></a>X     Expr; <span class="comment">// expression</span>
        <a id="L178"></a>Index Expr; <span class="comment">// index expression or beginning of slice range</span>
        <a id="L179"></a>End   Expr; <span class="comment">// end of slice range; or nil</span>
    <a id="L180"></a>};

    <a id="L182"></a><span class="comment">// A TypeAssertExpr node represents an expression followed by a</span>
    <a id="L183"></a><span class="comment">// type assertion.</span>
    <a id="L184"></a><span class="comment">//</span>
    <a id="L185"></a>TypeAssertExpr struct {
        <a id="L186"></a>X    Expr; <span class="comment">// expression</span>
        <a id="L187"></a>Type Expr; <span class="comment">// asserted type; nil means type switch X.(type)</span>
    <a id="L188"></a>};

    <a id="L190"></a><span class="comment">// A CallExpr node represents an expression followed by an argument list.</span>
    <a id="L191"></a>CallExpr struct {
        <a id="L192"></a>Fun    Expr;           <span class="comment">// function expression</span>
        <a id="L193"></a>Lparen token.Position; <span class="comment">// position of &#34;(&#34;</span>
        <a id="L194"></a>Args   []Expr;         <span class="comment">// function arguments</span>
        <a id="L195"></a>Rparen token.Position; <span class="comment">// positions of &#34;)&#34;</span>
    <a id="L196"></a>};

    <a id="L198"></a><span class="comment">// A StarExpr node represents an expression of the form &#34;*&#34; Expression.</span>
    <a id="L199"></a><span class="comment">// Semantically it could be a unary &#34;*&#34; expression, or a pointer type.</span>
    <a id="L200"></a>StarExpr struct {
        <a id="L201"></a>token.Position;       <span class="comment">// position of &#34;*&#34;</span>
        <a id="L202"></a>X               Expr; <span class="comment">// operand</span>
    <a id="L203"></a>};

    <a id="L205"></a><span class="comment">// A UnaryExpr node represents a unary expression.</span>
    <a id="L206"></a><span class="comment">// Unary &#34;*&#34; expressions are represented via StarExpr nodes.</span>
    <a id="L207"></a><span class="comment">//</span>
    <a id="L208"></a>UnaryExpr struct {
        <a id="L209"></a>token.Position;              <span class="comment">// position of Op</span>
        <a id="L210"></a>Op              token.Token; <span class="comment">// operator</span>
        <a id="L211"></a>X               Expr;        <span class="comment">// operand</span>
    <a id="L212"></a>};

    <a id="L214"></a><span class="comment">// A BinaryExpr node represents a binary expression.</span>
    <a id="L215"></a><span class="comment">//</span>
    <a id="L216"></a>BinaryExpr struct {
        <a id="L217"></a>X     Expr;           <span class="comment">// left operand</span>
        <a id="L218"></a>OpPos token.Position; <span class="comment">// position of Op</span>
        <a id="L219"></a>Op    token.Token;    <span class="comment">// operator</span>
        <a id="L220"></a>Y     Expr;           <span class="comment">// right operand</span>
    <a id="L221"></a>};

    <a id="L223"></a><span class="comment">// A KeyValueExpr node represents (key : value) pairs</span>
    <a id="L224"></a><span class="comment">// in composite literals.</span>
    <a id="L225"></a><span class="comment">//</span>
    <a id="L226"></a>KeyValueExpr struct {
        <a id="L227"></a>Key   Expr;
        <a id="L228"></a>Colon token.Position; <span class="comment">// position of &#34;:&#34;</span>
        <a id="L229"></a>Value Expr;
    <a id="L230"></a>};
<a id="L231"></a>)


<a id="L234"></a><span class="comment">// The direction of a channel type is indicated by one</span>
<a id="L235"></a><span class="comment">// of the following constants.</span>
<a id="L236"></a><span class="comment">//</span>
<a id="L237"></a>type ChanDir int

<a id="L239"></a>const (
    <a id="L240"></a>SEND ChanDir = 1 &lt;&lt; iota;
    <a id="L241"></a>RECV;
<a id="L242"></a>)


<a id="L245"></a><span class="comment">// A type is represented by a tree consisting of one</span>
<a id="L246"></a><span class="comment">// or more of the following type-specific expression</span>
<a id="L247"></a><span class="comment">// nodes.</span>
<a id="L248"></a><span class="comment">//</span>
<a id="L249"></a>type (
    <a id="L250"></a><span class="comment">// An ArrayType node represents an array or slice type.</span>
    <a id="L251"></a>ArrayType struct {
        <a id="L252"></a>token.Position;       <span class="comment">// position of &#34;[&#34;</span>
        <a id="L253"></a>Len             Expr; <span class="comment">// Ellipsis node for [...]T array types, nil for slice types</span>
        <a id="L254"></a>Elt             Expr; <span class="comment">// element type</span>
    <a id="L255"></a>};

    <a id="L257"></a><span class="comment">// A StructType node represents a struct type.</span>
    <a id="L258"></a>StructType struct {
        <a id="L259"></a>token.Position;                 <span class="comment">// position of &#34;struct&#34; keyword</span>
        <a id="L260"></a>Lbrace          token.Position; <span class="comment">// position of &#34;{&#34;</span>
        <a id="L261"></a>Fields          []*Field;       <span class="comment">// list of field declarations</span>
        <a id="L262"></a>Rbrace          token.Position; <span class="comment">// position of &#34;}&#34;</span>
        <a id="L263"></a>Incomplete      bool;           <span class="comment">// true if (source) fields are missing in the Fields list</span>
    <a id="L264"></a>};

    <a id="L266"></a><span class="comment">// Pointer types are represented via StarExpr nodes.</span>

    <a id="L268"></a><span class="comment">// A FuncType node represents a function type.</span>
    <a id="L269"></a>FuncType struct {
        <a id="L270"></a>token.Position;           <span class="comment">// position of &#34;func&#34; keyword</span>
        <a id="L271"></a>Params          []*Field; <span class="comment">// (incoming) parameters</span>
        <a id="L272"></a>Results         []*Field; <span class="comment">// (outgoing) results</span>
    <a id="L273"></a>};

    <a id="L275"></a><span class="comment">// An InterfaceType node represents an interface type.</span>
    <a id="L276"></a>InterfaceType struct {
        <a id="L277"></a>token.Position;                 <span class="comment">// position of &#34;interface&#34; keyword</span>
        <a id="L278"></a>Lbrace          token.Position; <span class="comment">// position of &#34;{&#34;</span>
        <a id="L279"></a>Methods         []*Field;       <span class="comment">// list of methods</span>
        <a id="L280"></a>Rbrace          token.Position; <span class="comment">// position of &#34;}&#34;</span>
        <a id="L281"></a>Incomplete      bool;           <span class="comment">// true if (source) methods are missing in the Methods list</span>
    <a id="L282"></a>};

    <a id="L284"></a><span class="comment">// A MapType node represents a map type.</span>
    <a id="L285"></a>MapType struct {
        <a id="L286"></a>token.Position; <span class="comment">// position of &#34;map&#34; keyword</span>
        <a id="L287"></a>Key             Expr;
        <a id="L288"></a>Value           Expr;
    <a id="L289"></a>};

    <a id="L291"></a><span class="comment">// A ChanType node represents a channel type.</span>
    <a id="L292"></a>ChanType struct {
        <a id="L293"></a>token.Position;          <span class="comment">// position of &#34;chan&#34; keyword or &#34;&lt;-&#34; (whichever comes first)</span>
        <a id="L294"></a>Dir             ChanDir; <span class="comment">// channel direction</span>
        <a id="L295"></a>Value           Expr;    <span class="comment">// value type</span>
    <a id="L296"></a>};
<a id="L297"></a>)


<a id="L300"></a><span class="comment">// Pos() implementations for expression/type where the position</span>
<a id="L301"></a><span class="comment">// corresponds to the position of a sub-node.</span>
<a id="L302"></a><span class="comment">//</span>
<a id="L303"></a>func (x *StringList) Pos() token.Position     { return x.Strings[0].Pos() }
<a id="L304"></a>func (x *FuncLit) Pos() token.Position        { return x.Type.Pos() }
<a id="L305"></a>func (x *CompositeLit) Pos() token.Position   { return x.Type.Pos() }
<a id="L306"></a>func (x *SelectorExpr) Pos() token.Position   { return x.X.Pos() }
<a id="L307"></a>func (x *IndexExpr) Pos() token.Position      { return x.X.Pos() }
<a id="L308"></a>func (x *TypeAssertExpr) Pos() token.Position { return x.X.Pos() }
<a id="L309"></a>func (x *CallExpr) Pos() token.Position       { return x.Fun.Pos() }
<a id="L310"></a>func (x *BinaryExpr) Pos() token.Position     { return x.X.Pos() }
<a id="L311"></a>func (x *KeyValueExpr) Pos() token.Position   { return x.Key.Pos() }


<a id="L314"></a><span class="comment">// exprNode() ensures that only expression/type nodes can be</span>
<a id="L315"></a><span class="comment">// assigned to an ExprNode.</span>
<a id="L316"></a>func (x *BadExpr) exprNode()        {}
<a id="L317"></a>func (x *Ident) exprNode()          {}
<a id="L318"></a>func (x *Ellipsis) exprNode()       {}
<a id="L319"></a>func (x *BasicLit) exprNode()       {}
<a id="L320"></a>func (x *StringList) exprNode()     {}
<a id="L321"></a>func (x *FuncLit) exprNode()        {}
<a id="L322"></a>func (x *CompositeLit) exprNode()   {}
<a id="L323"></a>func (x *ParenExpr) exprNode()      {}
<a id="L324"></a>func (x *SelectorExpr) exprNode()   {}
<a id="L325"></a>func (x *IndexExpr) exprNode()      {}
<a id="L326"></a>func (x *TypeAssertExpr) exprNode() {}
<a id="L327"></a>func (x *CallExpr) exprNode()       {}
<a id="L328"></a>func (x *StarExpr) exprNode()       {}
<a id="L329"></a>func (x *UnaryExpr) exprNode()      {}
<a id="L330"></a>func (x *BinaryExpr) exprNode()     {}
<a id="L331"></a>func (x *KeyValueExpr) exprNode()   {}

<a id="L333"></a>func (x *ArrayType) exprNode()     {}
<a id="L334"></a>func (x *StructType) exprNode()    {}
<a id="L335"></a>func (x *FuncType) exprNode()      {}
<a id="L336"></a>func (x *InterfaceType) exprNode() {}
<a id="L337"></a>func (x *MapType) exprNode()       {}
<a id="L338"></a>func (x *ChanType) exprNode()      {}


<a id="L341"></a><span class="comment">// IsExported returns whether name is an exported Go symbol</span>
<a id="L342"></a><span class="comment">// (i.e., whether it begins with an uppercase letter).</span>
<a id="L343"></a>func IsExported(name string) bool {
    <a id="L344"></a>ch, _ := utf8.DecodeRuneInString(name);
    <a id="L345"></a>return unicode.IsUpper(ch);
<a id="L346"></a>}

<a id="L348"></a><span class="comment">// IsExported returns whether name is an exported Go symbol</span>
<a id="L349"></a><span class="comment">// (i.e., whether it begins with an uppercase letter).</span>
<a id="L350"></a>func (name *Ident) IsExported() bool { return IsExported(name.Value) }

<a id="L352"></a>func (name *Ident) String() string { return name.Value }


<a id="L355"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L356"></a><span class="comment">// Statements</span>

<a id="L358"></a><span class="comment">// A statement is represented by a tree consisting of one</span>
<a id="L359"></a><span class="comment">// or more of the following concrete statement nodes.</span>
<a id="L360"></a><span class="comment">//</span>
<a id="L361"></a>type (
    <a id="L362"></a><span class="comment">// A BadStmt node is a placeholder for statements containing</span>
    <a id="L363"></a><span class="comment">// syntax errors for which no correct statement nodes can be</span>
    <a id="L364"></a><span class="comment">// created.</span>
    <a id="L365"></a><span class="comment">//</span>
    <a id="L366"></a>BadStmt struct {
        <a id="L367"></a>token.Position; <span class="comment">// beginning position of bad statement</span>
    <a id="L368"></a>};

    <a id="L370"></a><span class="comment">// A DeclStmt node represents a declaration in a statement list.</span>
    <a id="L371"></a>DeclStmt struct {
        <a id="L372"></a>Decl Decl;
    <a id="L373"></a>};

    <a id="L375"></a><span class="comment">// An EmptyStmt node represents an empty statement.</span>
    <a id="L376"></a><span class="comment">// The &#34;position&#34; of the empty statement is the position</span>
    <a id="L377"></a><span class="comment">// of the immediately preceeding semicolon.</span>
    <a id="L378"></a><span class="comment">//</span>
    <a id="L379"></a>EmptyStmt struct {
        <a id="L380"></a>token.Position; <span class="comment">// position of preceeding &#34;;&#34;</span>
    <a id="L381"></a>};

    <a id="L383"></a><span class="comment">// A LabeledStmt node represents a labeled statement.</span>
    <a id="L384"></a>LabeledStmt struct {
        <a id="L385"></a>Label *Ident;
        <a id="L386"></a>Stmt  Stmt;
    <a id="L387"></a>};

    <a id="L389"></a><span class="comment">// An ExprStmt node represents a (stand-alone) expression</span>
    <a id="L390"></a><span class="comment">// in a statement list.</span>
    <a id="L391"></a><span class="comment">//</span>
    <a id="L392"></a>ExprStmt struct {
        <a id="L393"></a>X Expr; <span class="comment">// expression</span>
    <a id="L394"></a>};

    <a id="L396"></a><span class="comment">// An IncDecStmt node represents an increment or decrement statement.</span>
    <a id="L397"></a>IncDecStmt struct {
        <a id="L398"></a>X   Expr;
        <a id="L399"></a>Tok token.Token; <span class="comment">// INC or DEC</span>
    <a id="L400"></a>};

    <a id="L402"></a><span class="comment">// An AssignStmt node represents an assignment or</span>
    <a id="L403"></a><span class="comment">// a short variable declaration.</span>
    <a id="L404"></a>AssignStmt struct {
        <a id="L405"></a>Lhs    []Expr;
        <a id="L406"></a>TokPos token.Position; <span class="comment">// position of Tok</span>
        <a id="L407"></a>Tok    token.Token;    <span class="comment">// assignment token, DEFINE</span>
        <a id="L408"></a>Rhs    []Expr;
    <a id="L409"></a>};

    <a id="L411"></a><span class="comment">// A GoStmt node represents a go statement.</span>
    <a id="L412"></a>GoStmt struct {
        <a id="L413"></a>token.Position; <span class="comment">// position of &#34;go&#34; keyword</span>
        <a id="L414"></a>Call            *CallExpr;
    <a id="L415"></a>};

    <a id="L417"></a><span class="comment">// A DeferStmt node represents a defer statement.</span>
    <a id="L418"></a>DeferStmt struct {
        <a id="L419"></a>token.Position; <span class="comment">// position of &#34;defer&#34; keyword</span>
        <a id="L420"></a>Call            *CallExpr;
    <a id="L421"></a>};

    <a id="L423"></a><span class="comment">// A ReturnStmt node represents a return statement.</span>
    <a id="L424"></a>ReturnStmt struct {
        <a id="L425"></a>token.Position; <span class="comment">// position of &#34;return&#34; keyword</span>
        <a id="L426"></a>Results         []Expr;
    <a id="L427"></a>};

    <a id="L429"></a><span class="comment">// A BranchStmt node represents a break, continue, goto,</span>
    <a id="L430"></a><span class="comment">// or fallthrough statement.</span>
    <a id="L431"></a><span class="comment">//</span>
    <a id="L432"></a>BranchStmt struct {
        <a id="L433"></a>token.Position;              <span class="comment">// position of Tok</span>
        <a id="L434"></a>Tok             token.Token; <span class="comment">// keyword token (BREAK, CONTINUE, GOTO, FALLTHROUGH)</span>
        <a id="L435"></a>Label           *Ident;
    <a id="L436"></a>};

    <a id="L438"></a><span class="comment">// A BlockStmt node represents a braced statement list.</span>
    <a id="L439"></a>BlockStmt struct {
        <a id="L440"></a>token.Position; <span class="comment">// position of &#34;{&#34;</span>
        <a id="L441"></a>List            []Stmt;
        <a id="L442"></a>Rbrace          token.Position; <span class="comment">// position of &#34;}&#34;</span>
    <a id="L443"></a>};

    <a id="L445"></a><span class="comment">// An IfStmt node represents an if statement.</span>
    <a id="L446"></a>IfStmt struct {
        <a id="L447"></a>token.Position; <span class="comment">// position of &#34;if&#34; keyword</span>
        <a id="L448"></a>Init            Stmt;
        <a id="L449"></a>Cond            Expr;
        <a id="L450"></a>Body            *BlockStmt;
        <a id="L451"></a>Else            Stmt;
    <a id="L452"></a>};

    <a id="L454"></a><span class="comment">// A CaseClause represents a case of an expression switch statement.</span>
    <a id="L455"></a>CaseClause struct {
        <a id="L456"></a>token.Position;                 <span class="comment">// position of &#34;case&#34; or &#34;default&#34; keyword</span>
        <a id="L457"></a>Values          []Expr;         <span class="comment">// nil means default case</span>
        <a id="L458"></a>Colon           token.Position; <span class="comment">// position of &#34;:&#34;</span>
        <a id="L459"></a>Body            []Stmt;         <span class="comment">// statement list; or nil</span>
    <a id="L460"></a>};

    <a id="L462"></a><span class="comment">// A SwitchStmt node represents an expression switch statement.</span>
    <a id="L463"></a>SwitchStmt struct {
        <a id="L464"></a>token.Position; <span class="comment">// position of &#34;switch&#34; keyword</span>
        <a id="L465"></a>Init            Stmt;
        <a id="L466"></a>Tag             Expr;
        <a id="L467"></a>Body            *BlockStmt; <span class="comment">// CaseClauses only</span>
    <a id="L468"></a>};

    <a id="L470"></a><span class="comment">// A TypeCaseClause represents a case of a type switch statement.</span>
    <a id="L471"></a>TypeCaseClause struct {
        <a id="L472"></a>token.Position;                 <span class="comment">// position of &#34;case&#34; or &#34;default&#34; keyword</span>
        <a id="L473"></a>Types           []Expr;         <span class="comment">// nil means default case</span>
        <a id="L474"></a>Colon           token.Position; <span class="comment">// position of &#34;:&#34;</span>
        <a id="L475"></a>Body            []Stmt;         <span class="comment">// statement list; or nil</span>
    <a id="L476"></a>};

    <a id="L478"></a><span class="comment">// An TypeSwitchStmt node represents a type switch statement.</span>
    <a id="L479"></a>TypeSwitchStmt struct {
        <a id="L480"></a>token.Position; <span class="comment">// position of &#34;switch&#34; keyword</span>
        <a id="L481"></a>Init            Stmt;
        <a id="L482"></a>Assign          Stmt;       <span class="comment">// x := y.(type)</span>
        <a id="L483"></a>Body            *BlockStmt; <span class="comment">// TypeCaseClauses only</span>
    <a id="L484"></a>};

    <a id="L486"></a><span class="comment">// A CommClause node represents a case of a select statement.</span>
    <a id="L487"></a>CommClause struct {
        <a id="L488"></a>token.Position;                 <span class="comment">// position of &#34;case&#34; or &#34;default&#34; keyword</span>
        <a id="L489"></a>Tok             token.Token;    <span class="comment">// ASSIGN or DEFINE (valid only if Lhs != nil)</span>
        <a id="L490"></a>Lhs, Rhs        Expr;           <span class="comment">// Rhs == nil means default case</span>
        <a id="L491"></a>Colon           token.Position; <span class="comment">// position of &#34;:&#34;</span>
        <a id="L492"></a>Body            []Stmt;         <span class="comment">// statement list; or nil</span>
    <a id="L493"></a>};

    <a id="L495"></a><span class="comment">// An SelectStmt node represents a select statement.</span>
    <a id="L496"></a>SelectStmt struct {
        <a id="L497"></a>token.Position;             <span class="comment">// position of &#34;select&#34; keyword</span>
        <a id="L498"></a>Body            *BlockStmt; <span class="comment">// CommClauses only</span>
    <a id="L499"></a>};

    <a id="L501"></a><span class="comment">// A ForStmt represents a for statement.</span>
    <a id="L502"></a>ForStmt struct {
        <a id="L503"></a>token.Position; <span class="comment">// position of &#34;for&#34; keyword</span>
        <a id="L504"></a>Init            Stmt;
        <a id="L505"></a>Cond            Expr;
        <a id="L506"></a>Post            Stmt;
        <a id="L507"></a>Body            *BlockStmt;
    <a id="L508"></a>};

    <a id="L510"></a><span class="comment">// A RangeStmt represents a for statement with a range clause.</span>
    <a id="L511"></a>RangeStmt struct {
        <a id="L512"></a>token.Position;                 <span class="comment">// position of &#34;for&#34; keyword</span>
        <a id="L513"></a>Key, Value      Expr;           <span class="comment">// Value may be nil</span>
        <a id="L514"></a>TokPos          token.Position; <span class="comment">// position of Tok</span>
        <a id="L515"></a>Tok             token.Token;    <span class="comment">// ASSIGN, DEFINE</span>
        <a id="L516"></a>X               Expr;           <span class="comment">// value to range over</span>
        <a id="L517"></a>Body            *BlockStmt;
    <a id="L518"></a>};
<a id="L519"></a>)


<a id="L522"></a><span class="comment">// Pos() implementations for statement nodes where the position</span>
<a id="L523"></a><span class="comment">// corresponds to the position of a sub-node.</span>
<a id="L524"></a><span class="comment">//</span>
<a id="L525"></a>func (s *DeclStmt) Pos() token.Position    { return s.Decl.Pos() }
<a id="L526"></a>func (s *LabeledStmt) Pos() token.Position { return s.Label.Pos() }
<a id="L527"></a>func (s *ExprStmt) Pos() token.Position    { return s.X.Pos() }
<a id="L528"></a>func (s *IncDecStmt) Pos() token.Position  { return s.X.Pos() }
<a id="L529"></a>func (s *AssignStmt) Pos() token.Position  { return s.Lhs[0].Pos() }


<a id="L532"></a><span class="comment">// stmtNode() ensures that only statement nodes can be</span>
<a id="L533"></a><span class="comment">// assigned to a StmtNode.</span>
<a id="L534"></a><span class="comment">//</span>
<a id="L535"></a>func (s *BadStmt) stmtNode()        {}
<a id="L536"></a>func (s *DeclStmt) stmtNode()       {}
<a id="L537"></a>func (s *EmptyStmt) stmtNode()      {}
<a id="L538"></a>func (s *LabeledStmt) stmtNode()    {}
<a id="L539"></a>func (s *ExprStmt) stmtNode()       {}
<a id="L540"></a>func (s *IncDecStmt) stmtNode()     {}
<a id="L541"></a>func (s *AssignStmt) stmtNode()     {}
<a id="L542"></a>func (s *GoStmt) stmtNode()         {}
<a id="L543"></a>func (s *DeferStmt) stmtNode()      {}
<a id="L544"></a>func (s *ReturnStmt) stmtNode()     {}
<a id="L545"></a>func (s *BranchStmt) stmtNode()     {}
<a id="L546"></a>func (s *BlockStmt) stmtNode()      {}
<a id="L547"></a>func (s *IfStmt) stmtNode()         {}
<a id="L548"></a>func (s *CaseClause) stmtNode()     {}
<a id="L549"></a>func (s *SwitchStmt) stmtNode()     {}
<a id="L550"></a>func (s *TypeCaseClause) stmtNode() {}
<a id="L551"></a>func (s *TypeSwitchStmt) stmtNode() {}
<a id="L552"></a>func (s *CommClause) stmtNode()     {}
<a id="L553"></a>func (s *SelectStmt) stmtNode()     {}
<a id="L554"></a>func (s *ForStmt) stmtNode()        {}
<a id="L555"></a>func (s *RangeStmt) stmtNode()      {}


<a id="L558"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L559"></a><span class="comment">// Declarations</span>

<a id="L561"></a><span class="comment">// A Spec node represents a single (non-parenthesized) import,</span>
<a id="L562"></a><span class="comment">// constant, type, or variable declaration.</span>
<a id="L563"></a><span class="comment">//</span>
<a id="L564"></a>type (
    <a id="L565"></a><span class="comment">// The Spec type stands for any of *ImportSpec, *ValueSpec, and *TypeSpec.</span>
    <a id="L566"></a>Spec interface {
        <a id="L567"></a>Node;
        <a id="L568"></a>specNode();
    <a id="L569"></a>};

    <a id="L571"></a><span class="comment">// An ImportSpec node represents a single package import.</span>
    <a id="L572"></a>ImportSpec struct {
        <a id="L573"></a>Doc     *CommentGroup; <span class="comment">// associated documentation; or nil</span>
        <a id="L574"></a>Name    *Ident;        <span class="comment">// local package name (including &#34;.&#34;); or nil</span>
        <a id="L575"></a>Path    []*BasicLit;   <span class="comment">// package path</span>
        <a id="L576"></a>Comment *CommentGroup; <span class="comment">// line comments; or nil</span>
    <a id="L577"></a>};

    <a id="L579"></a><span class="comment">// A ValueSpec node represents a constant or variable declaration</span>
    <a id="L580"></a><span class="comment">// (ConstSpec or VarSpec production).</span>
    <a id="L581"></a>ValueSpec struct {
        <a id="L582"></a>Doc     *CommentGroup; <span class="comment">// associated documentation; or nil</span>
        <a id="L583"></a>Names   []*Ident;      <span class="comment">// value names</span>
        <a id="L584"></a>Type    Expr;          <span class="comment">// value type; or nil</span>
        <a id="L585"></a>Values  []Expr;        <span class="comment">// initial values; or nil</span>
        <a id="L586"></a>Comment *CommentGroup; <span class="comment">// line comments; or nil</span>
    <a id="L587"></a>};

    <a id="L589"></a><span class="comment">// A TypeSpec node represents a type declaration (TypeSpec production).</span>
    <a id="L590"></a>TypeSpec struct {
        <a id="L591"></a>Doc     *CommentGroup; <span class="comment">// associated documentation; or nil</span>
        <a id="L592"></a>Name    *Ident;        <span class="comment">// type name</span>
        <a id="L593"></a>Type    Expr;
        <a id="L594"></a>Comment *CommentGroup; <span class="comment">// line comments; or nil</span>
    <a id="L595"></a>};
<a id="L596"></a>)


<a id="L599"></a><span class="comment">// Pos() implementations for spec nodes.</span>
<a id="L600"></a><span class="comment">//</span>
<a id="L601"></a>func (s *ImportSpec) Pos() token.Position {
    <a id="L602"></a>if s.Name != nil {
        <a id="L603"></a>return s.Name.Pos()
    <a id="L604"></a>}
    <a id="L605"></a>return s.Path[0].Pos();
<a id="L606"></a>}

<a id="L608"></a>func (s *ValueSpec) Pos() token.Position { return s.Names[0].Pos() }
<a id="L609"></a>func (s *TypeSpec) Pos() token.Position  { return s.Name.Pos() }


<a id="L612"></a><span class="comment">// specNode() ensures that only spec nodes can be</span>
<a id="L613"></a><span class="comment">// assigned to a Spec.</span>
<a id="L614"></a><span class="comment">//</span>
<a id="L615"></a>func (s *ImportSpec) specNode() {}
<a id="L616"></a>func (s *ValueSpec) specNode()  {}
<a id="L617"></a>func (s *TypeSpec) specNode()   {}


<a id="L620"></a><span class="comment">// A declaration is represented by one of the following declaration nodes.</span>
<a id="L621"></a><span class="comment">//</span>
<a id="L622"></a>type (
    <a id="L623"></a><span class="comment">// A BadDecl node is a placeholder for declarations containing</span>
    <a id="L624"></a><span class="comment">// syntax errors for which no correct declaration nodes can be</span>
    <a id="L625"></a><span class="comment">// created.</span>
    <a id="L626"></a><span class="comment">//</span>
    <a id="L627"></a>BadDecl struct {
        <a id="L628"></a>token.Position; <span class="comment">// beginning position of bad declaration</span>
    <a id="L629"></a>};

    <a id="L631"></a><span class="comment">// A GenDecl node (generic declaration node) represents an import,</span>
    <a id="L632"></a><span class="comment">// constant, type or variable declaration. A valid Lparen position</span>
    <a id="L633"></a><span class="comment">// (Lparen.Line &gt; 0) indicates a parenthesized declaration.</span>
    <a id="L634"></a><span class="comment">//</span>
    <a id="L635"></a><span class="comment">// Relationship between Tok value and Specs element type:</span>
    <a id="L636"></a><span class="comment">//</span>
    <a id="L637"></a><span class="comment">//	token.IMPORT  *ImportSpec</span>
    <a id="L638"></a><span class="comment">//	token.CONST   *ValueSpec</span>
    <a id="L639"></a><span class="comment">//	token.TYPE    *TypeSpec</span>
    <a id="L640"></a><span class="comment">//	token.VAR     *ValueSpec</span>
    <a id="L641"></a><span class="comment">//</span>
    <a id="L642"></a>GenDecl struct {
        <a id="L643"></a>Doc             *CommentGroup;  <span class="comment">// associated documentation; or nil</span>
        <a id="L644"></a>token.Position;                 <span class="comment">// position of Tok</span>
        <a id="L645"></a>Tok             token.Token;    <span class="comment">// IMPORT, CONST, TYPE, VAR</span>
        <a id="L646"></a>Lparen          token.Position; <span class="comment">// position of &#39;(&#39;, if any</span>
        <a id="L647"></a>Specs           []Spec;
        <a id="L648"></a>Rparen          token.Position; <span class="comment">// position of &#39;)&#39;, if any</span>
    <a id="L649"></a>};

    <a id="L651"></a><span class="comment">// A FuncDecl node represents a function declaration.</span>
    <a id="L652"></a>FuncDecl struct {
        <a id="L653"></a>Doc  *CommentGroup; <span class="comment">// associated documentation; or nil</span>
        <a id="L654"></a>Recv *Field;        <span class="comment">// receiver (methods); or nil (functions)</span>
        <a id="L655"></a>Name *Ident;        <span class="comment">// function/method name</span>
        <a id="L656"></a>Type *FuncType;     <span class="comment">// position of Func keyword, parameters and results</span>
        <a id="L657"></a>Body *BlockStmt;    <span class="comment">// function body; or nil (forward declaration)</span>
    <a id="L658"></a>};
<a id="L659"></a>)


<a id="L662"></a><span class="comment">// The position of a FuncDecl node is the position of its function type.</span>
<a id="L663"></a>func (d *FuncDecl) Pos() token.Position { return d.Type.Pos() }


<a id="L666"></a><span class="comment">// declNode() ensures that only declaration nodes can be</span>
<a id="L667"></a><span class="comment">// assigned to a DeclNode.</span>
<a id="L668"></a><span class="comment">//</span>
<a id="L669"></a>func (d *BadDecl) declNode()  {}
<a id="L670"></a>func (d *GenDecl) declNode()  {}
<a id="L671"></a>func (d *FuncDecl) declNode() {}


<a id="L674"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L675"></a><span class="comment">// Files and packages</span>

<a id="L677"></a><span class="comment">// A File node represents a Go source file.</span>
<a id="L678"></a><span class="comment">//</span>
<a id="L679"></a>type File struct {
    <a id="L680"></a>Doc             *CommentGroup; <span class="comment">// associated documentation; or nil</span>
    <a id="L681"></a>token.Position;                <span class="comment">// position of &#34;package&#34; keyword</span>
    <a id="L682"></a>Name            *Ident;        <span class="comment">// package name</span>
    <a id="L683"></a>Decls           []Decl;        <span class="comment">// top-level declarations</span>
    <a id="L684"></a>Comments        *CommentGroup; <span class="comment">// list of all comments in the source file</span>
<a id="L685"></a>}


<a id="L688"></a><span class="comment">// A Package node represents a set of source files</span>
<a id="L689"></a><span class="comment">// collectively building a Go package.</span>
<a id="L690"></a><span class="comment">//</span>
<a id="L691"></a>type Package struct {
    <a id="L692"></a>Name  string;           <span class="comment">// package name</span>
    <a id="L693"></a>Path  string;           <span class="comment">// package path</span>
    <a id="L694"></a>Files map[string]*File; <span class="comment">// path-relative filenames</span>
<a id="L695"></a>}
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
