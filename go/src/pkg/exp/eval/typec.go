<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/eval/typec.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/exp/eval/typec.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package eval

<a id="L7"></a>import (
    <a id="L8"></a>&#34;go/ast&#34;;
    <a id="L9"></a>&#34;go/token&#34;;
    <a id="L10"></a>&#34;log&#34;;
<a id="L11"></a>)


<a id="L14"></a><span class="comment">/*</span>
<a id="L15"></a><span class="comment"> * Type compiler</span>
<a id="L16"></a><span class="comment"> */</span>

<a id="L18"></a>type typeCompiler struct {
    <a id="L19"></a>*compiler;
    <a id="L20"></a>block *block;
    <a id="L21"></a><span class="comment">// Check to be performed after a type declaration is compiled.</span>
    <a id="L22"></a><span class="comment">//</span>
    <a id="L23"></a><span class="comment">// TODO(austin) This will probably have to change after we</span>
    <a id="L24"></a><span class="comment">// eliminate forward declarations.</span>
    <a id="L25"></a>lateCheck func() bool;
<a id="L26"></a>}

<a id="L28"></a>func (a *typeCompiler) compileIdent(x *ast.Ident, allowRec bool) Type {
    <a id="L29"></a>_, _, def := a.block.Lookup(x.Value);
    <a id="L30"></a>if def == nil {
        <a id="L31"></a>a.diagAt(x, &#34;%s: undefined&#34;, x.Value);
        <a id="L32"></a>return nil;
    <a id="L33"></a>}
    <a id="L34"></a>switch def := def.(type) {
    <a id="L35"></a>case *Constant:
        <a id="L36"></a>a.diagAt(x, &#34;constant %v used as type&#34;, x.Value);
        <a id="L37"></a>return nil;
    <a id="L38"></a>case *Variable:
        <a id="L39"></a>a.diagAt(x, &#34;variable %v used as type&#34;, x.Value);
        <a id="L40"></a>return nil;
    <a id="L41"></a>case *NamedType:
        <a id="L42"></a>if !allowRec &amp;&amp; def.incomplete {
            <a id="L43"></a>a.diagAt(x, &#34;illegal recursive type&#34;);
            <a id="L44"></a>return nil;
        <a id="L45"></a>}
        <a id="L46"></a>if !def.incomplete &amp;&amp; def.Def == nil {
            <a id="L47"></a><span class="comment">// Placeholder type from an earlier error</span>
            <a id="L48"></a>return nil
        <a id="L49"></a>}
        <a id="L50"></a>return def;
    <a id="L51"></a>case Type:
        <a id="L52"></a>return def
    <a id="L53"></a>}
    <a id="L54"></a>log.Crashf(&#34;name %s has unknown type %T&#34;, x.Value, def);
    <a id="L55"></a>return nil;
<a id="L56"></a>}

<a id="L58"></a>func (a *typeCompiler) compileArrayType(x *ast.ArrayType, allowRec bool) Type {
    <a id="L59"></a><span class="comment">// Compile element type</span>
    <a id="L60"></a>elem := a.compileType(x.Elt, allowRec);

    <a id="L62"></a><span class="comment">// Compile length expression</span>
    <a id="L63"></a>if x.Len == nil {
        <a id="L64"></a>if elem == nil {
            <a id="L65"></a>return nil
        <a id="L66"></a>}
        <a id="L67"></a>return NewSliceType(elem);
    <a id="L68"></a>}

    <a id="L70"></a>if _, ok := x.Len.(*ast.Ellipsis); ok {
        <a id="L71"></a>a.diagAt(x.Len, &#34;... array initailizers not implemented&#34;);
        <a id="L72"></a>return nil;
    <a id="L73"></a>}
    <a id="L74"></a>l, ok := a.compileArrayLen(a.block, x.Len);
    <a id="L75"></a>if !ok {
        <a id="L76"></a>return nil
    <a id="L77"></a>}
    <a id="L78"></a>if l &lt; 0 {
        <a id="L79"></a>a.diagAt(x.Len, &#34;array length must be non-negative&#34;);
        <a id="L80"></a>return nil;
    <a id="L81"></a>}
    <a id="L82"></a>if elem == nil {
        <a id="L83"></a>return nil
    <a id="L84"></a>}

    <a id="L86"></a>return NewArrayType(l, elem);
<a id="L87"></a>}

<a id="L89"></a>func countFields(fs []*ast.Field) int {
    <a id="L90"></a>n := 0;
    <a id="L91"></a>for _, f := range fs {
        <a id="L92"></a>if f.Names == nil {
            <a id="L93"></a>n++
        <a id="L94"></a>} else {
            <a id="L95"></a>n += len(f.Names)
        <a id="L96"></a>}
    <a id="L97"></a>}
    <a id="L98"></a>return n;
<a id="L99"></a>}

<a id="L101"></a>func (a *typeCompiler) compileFields(fs []*ast.Field, allowRec bool) ([]Type, []*ast.Ident, []token.Position, bool) {
    <a id="L102"></a>n := countFields(fs);
    <a id="L103"></a>ts := make([]Type, n);
    <a id="L104"></a>ns := make([]*ast.Ident, n);
    <a id="L105"></a>ps := make([]token.Position, n);

    <a id="L107"></a>bad := false;
    <a id="L108"></a>i := 0;
    <a id="L109"></a>for _, f := range fs {
        <a id="L110"></a>t := a.compileType(f.Type, allowRec);
        <a id="L111"></a>if t == nil {
            <a id="L112"></a>bad = true
        <a id="L113"></a>}
        <a id="L114"></a>if f.Names == nil {
            <a id="L115"></a>ns[i] = nil;
            <a id="L116"></a>ts[i] = t;
            <a id="L117"></a>ps[i] = f.Type.Pos();
            <a id="L118"></a>i++;
            <a id="L119"></a>continue;
        <a id="L120"></a>}
        <a id="L121"></a>for _, n := range f.Names {
            <a id="L122"></a>ns[i] = n;
            <a id="L123"></a>ts[i] = t;
            <a id="L124"></a>ps[i] = n.Pos();
            <a id="L125"></a>i++;
        <a id="L126"></a>}
    <a id="L127"></a>}

    <a id="L129"></a>return ts, ns, ps, bad;
<a id="L130"></a>}

<a id="L132"></a>func (a *typeCompiler) compileStructType(x *ast.StructType, allowRec bool) Type {
    <a id="L133"></a>ts, names, poss, bad := a.compileFields(x.Fields, allowRec);

    <a id="L135"></a><span class="comment">// XXX(Spec) The spec claims that field identifiers must be</span>
    <a id="L136"></a><span class="comment">// unique, but 6g only checks this when they are accessed.  I</span>
    <a id="L137"></a><span class="comment">// think the spec is better in this regard: if I write two</span>
    <a id="L138"></a><span class="comment">// fields with the same name in the same struct type, clearly</span>
    <a id="L139"></a><span class="comment">// that&#39;s a mistake.  This definition does *not* descend into</span>
    <a id="L140"></a><span class="comment">// anonymous fields, so it doesn&#39;t matter if those change.</span>
    <a id="L141"></a><span class="comment">// There&#39;s separate language in the spec about checking</span>
    <a id="L142"></a><span class="comment">// uniqueness of field names inherited from anonymous fields</span>
    <a id="L143"></a><span class="comment">// at use time.</span>
    <a id="L144"></a>fields := make([]StructField, len(ts));
    <a id="L145"></a>nameSet := make(map[string]token.Position, len(ts));
    <a id="L146"></a>for i := range fields {
        <a id="L147"></a><span class="comment">// Compute field name and check anonymous fields</span>
        <a id="L148"></a>var name string;
        <a id="L149"></a>if names[i] != nil {
            <a id="L150"></a>name = names[i].Value
        <a id="L151"></a>} else {
            <a id="L152"></a>if ts[i] == nil {
                <a id="L153"></a>continue
            <a id="L154"></a>}

            <a id="L156"></a>var nt *NamedType;
            <a id="L157"></a><span class="comment">// [For anonymous fields,] the unqualified</span>
            <a id="L158"></a><span class="comment">// type name acts as the field identifier.</span>
            <a id="L159"></a>switch t := ts[i].(type) {
            <a id="L160"></a>case *NamedType:
                <a id="L161"></a>name = t.Name;
                <a id="L162"></a>nt = t;
            <a id="L163"></a>case *PtrType:
                <a id="L164"></a>switch t := t.Elem.(type) {
                <a id="L165"></a>case *NamedType:
                    <a id="L166"></a>name = t.Name;
                    <a id="L167"></a>nt = t;
                <a id="L168"></a>}
            <a id="L169"></a>}
            <a id="L170"></a><span class="comment">// [An anonymous field] must be specified as a</span>
            <a id="L171"></a><span class="comment">// type name T or as a pointer to a type name</span>
            <a id="L172"></a><span class="comment">// *T, and T itself, may not be a pointer or</span>
            <a id="L173"></a><span class="comment">// interface type.</span>
            <a id="L174"></a>if nt == nil {
                <a id="L175"></a>a.diagAt(&amp;poss[i], &#34;embedded type must T or *T, where T is a named type&#34;);
                <a id="L176"></a>bad = true;
                <a id="L177"></a>continue;
            <a id="L178"></a>}
            <a id="L179"></a><span class="comment">// The check for embedded pointer types must</span>
            <a id="L180"></a><span class="comment">// be deferred because of things like</span>
            <a id="L181"></a><span class="comment">//  type T *struct { T }</span>
            <a id="L182"></a>lateCheck := a.lateCheck;
            <a id="L183"></a>a.lateCheck = func() bool {
                <a id="L184"></a>if _, ok := nt.lit().(*PtrType); ok {
                    <a id="L185"></a>a.diagAt(&amp;poss[i], &#34;embedded type %v is a pointer type&#34;, nt);
                    <a id="L186"></a>return false;
                <a id="L187"></a>}
                <a id="L188"></a>return lateCheck();
            <a id="L189"></a>};
        <a id="L190"></a>}

        <a id="L192"></a><span class="comment">// Check name uniqueness</span>
        <a id="L193"></a>if prev, ok := nameSet[name]; ok {
            <a id="L194"></a>a.diagAt(&amp;poss[i], &#34;field %s redeclared\n\tprevious declaration at %s&#34;, name, &amp;prev);
            <a id="L195"></a>bad = true;
            <a id="L196"></a>continue;
        <a id="L197"></a>}
        <a id="L198"></a>nameSet[name] = poss[i];

        <a id="L200"></a><span class="comment">// Create field</span>
        <a id="L201"></a>fields[i].Name = name;
        <a id="L202"></a>fields[i].Type = ts[i];
        <a id="L203"></a>fields[i].Anonymous = (names[i] == nil);
    <a id="L204"></a>}

    <a id="L206"></a>if bad {
        <a id="L207"></a>return nil
    <a id="L208"></a>}

    <a id="L210"></a>return NewStructType(fields);
<a id="L211"></a>}

<a id="L213"></a>func (a *typeCompiler) compilePtrType(x *ast.StarExpr) Type {
    <a id="L214"></a>elem := a.compileType(x.X, true);
    <a id="L215"></a>if elem == nil {
        <a id="L216"></a>return nil
    <a id="L217"></a>}
    <a id="L218"></a>return NewPtrType(elem);
<a id="L219"></a>}

<a id="L221"></a>func (a *typeCompiler) compileFuncType(x *ast.FuncType, allowRec bool) *FuncDecl {
    <a id="L222"></a><span class="comment">// TODO(austin) Variadic function types</span>

    <a id="L224"></a><span class="comment">// The types of parameters and results must be complete.</span>
    <a id="L225"></a><span class="comment">//</span>
    <a id="L226"></a><span class="comment">// TODO(austin) It&#39;s not clear they actually have to be complete.</span>
    <a id="L227"></a>in, inNames, _, inBad := a.compileFields(x.Params, allowRec);
    <a id="L228"></a>out, outNames, _, outBad := a.compileFields(x.Results, allowRec);

    <a id="L230"></a>if inBad || outBad {
        <a id="L231"></a>return nil
    <a id="L232"></a>}
    <a id="L233"></a>return &amp;FuncDecl{NewFuncType(in, false, out), nil, inNames, outNames};
<a id="L234"></a>}

<a id="L236"></a>func (a *typeCompiler) compileInterfaceType(x *ast.InterfaceType, allowRec bool) *InterfaceType {
    <a id="L237"></a>ts, names, poss, bad := a.compileFields(x.Methods, allowRec);

    <a id="L239"></a>methods := make([]IMethod, len(ts));
    <a id="L240"></a>nameSet := make(map[string]token.Position, len(ts));
    <a id="L241"></a>embeds := make([]*InterfaceType, len(ts));

    <a id="L243"></a>var nm, ne int;
    <a id="L244"></a>for i := range ts {
        <a id="L245"></a>if ts[i] == nil {
            <a id="L246"></a>continue
        <a id="L247"></a>}

        <a id="L249"></a>if names[i] != nil {
            <a id="L250"></a>name := names[i].Value;
            <a id="L251"></a>methods[nm].Name = name;
            <a id="L252"></a>methods[nm].Type = ts[i].(*FuncType);
            <a id="L253"></a>nm++;
            <a id="L254"></a>if prev, ok := nameSet[name]; ok {
                <a id="L255"></a>a.diagAt(&amp;poss[i], &#34;method %s redeclared\n\tprevious declaration at %s&#34;, name, &amp;prev);
                <a id="L256"></a>bad = true;
                <a id="L257"></a>continue;
            <a id="L258"></a>}
            <a id="L259"></a>nameSet[name] = poss[i];
        <a id="L260"></a>} else {
            <a id="L261"></a><span class="comment">// Embedded interface</span>
            <a id="L262"></a>it, ok := ts[i].lit().(*InterfaceType);
            <a id="L263"></a>if !ok {
                <a id="L264"></a>a.diagAt(&amp;poss[i], &#34;embedded type must be an interface&#34;);
                <a id="L265"></a>bad = true;
                <a id="L266"></a>continue;
            <a id="L267"></a>}
            <a id="L268"></a>embeds[ne] = it;
            <a id="L269"></a>ne++;
            <a id="L270"></a>for _, m := range it.methods {
                <a id="L271"></a>if prev, ok := nameSet[m.Name]; ok {
                    <a id="L272"></a>a.diagAt(&amp;poss[i], &#34;method %s redeclared\n\tprevious declaration at %s&#34;, m.Name, &amp;prev);
                    <a id="L273"></a>bad = true;
                    <a id="L274"></a>continue;
                <a id="L275"></a>}
                <a id="L276"></a>nameSet[m.Name] = poss[i];
            <a id="L277"></a>}
        <a id="L278"></a>}
    <a id="L279"></a>}

    <a id="L281"></a>if bad {
        <a id="L282"></a>return nil
    <a id="L283"></a>}

    <a id="L285"></a>methods = methods[0:nm];
    <a id="L286"></a>embeds = embeds[0:ne];

    <a id="L288"></a>return NewInterfaceType(methods, embeds);
<a id="L289"></a>}

<a id="L291"></a>func (a *typeCompiler) compileMapType(x *ast.MapType) Type {
    <a id="L292"></a>key := a.compileType(x.Key, true);
    <a id="L293"></a>val := a.compileType(x.Value, true);
    <a id="L294"></a>if key == nil || val == nil {
        <a id="L295"></a>return nil
    <a id="L296"></a>}
    <a id="L297"></a><span class="comment">// XXX(Spec) The Map types section explicitly lists all types</span>
    <a id="L298"></a><span class="comment">// that can be map keys except for function types.</span>
    <a id="L299"></a>switch key.lit().(type) {
    <a id="L300"></a>case *StructType:
        <a id="L301"></a>a.diagAt(x, &#34;map key cannot be a struct type&#34;);
        <a id="L302"></a>return nil;
    <a id="L303"></a>case *ArrayType:
        <a id="L304"></a>a.diagAt(x, &#34;map key cannot be an array type&#34;);
        <a id="L305"></a>return nil;
    <a id="L306"></a>case *SliceType:
        <a id="L307"></a>a.diagAt(x, &#34;map key cannot be a slice type&#34;);
        <a id="L308"></a>return nil;
    <a id="L309"></a>}
    <a id="L310"></a>return NewMapType(key, val);
<a id="L311"></a>}

<a id="L313"></a>func (a *typeCompiler) compileType(x ast.Expr, allowRec bool) Type {
    <a id="L314"></a>switch x := x.(type) {
    <a id="L315"></a>case *ast.BadExpr:
        <a id="L316"></a><span class="comment">// Error already reported by parser</span>
        <a id="L317"></a>a.silentErrors++;
        <a id="L318"></a>return nil;

    <a id="L320"></a>case *ast.Ident:
        <a id="L321"></a>return a.compileIdent(x, allowRec)

    <a id="L323"></a>case *ast.ArrayType:
        <a id="L324"></a>return a.compileArrayType(x, allowRec)

    <a id="L326"></a>case *ast.StructType:
        <a id="L327"></a>return a.compileStructType(x, allowRec)

    <a id="L329"></a>case *ast.StarExpr:
        <a id="L330"></a>return a.compilePtrType(x)

    <a id="L332"></a>case *ast.FuncType:
        <a id="L333"></a>fd := a.compileFuncType(x, allowRec);
        <a id="L334"></a>if fd == nil {
            <a id="L335"></a>return nil
        <a id="L336"></a>}
        <a id="L337"></a>return fd.Type;

    <a id="L339"></a>case *ast.InterfaceType:
        <a id="L340"></a>return a.compileInterfaceType(x, allowRec)

    <a id="L342"></a>case *ast.MapType:
        <a id="L343"></a>return a.compileMapType(x)

    <a id="L345"></a>case *ast.ChanType:
        <a id="L346"></a>goto notimpl

    <a id="L348"></a>case *ast.ParenExpr:
        <a id="L349"></a>return a.compileType(x.X, allowRec)

    <a id="L351"></a>case *ast.Ellipsis:
        <a id="L352"></a>a.diagAt(x, &#34;illegal use of ellipsis&#34;);
        <a id="L353"></a>return nil;
    <a id="L354"></a>}
    <a id="L355"></a>a.diagAt(x, &#34;expression used as type&#34;);
    <a id="L356"></a>return nil;

<a id="L358"></a>notimpl:
    <a id="L359"></a>a.diagAt(x, &#34;compileType: %T not implemented&#34;, x);
    <a id="L360"></a>return nil;
<a id="L361"></a>}

<a id="L363"></a><span class="comment">/*</span>
<a id="L364"></a><span class="comment"> * Type compiler interface</span>
<a id="L365"></a><span class="comment"> */</span>

<a id="L367"></a>func noLateCheck() bool { return true }

<a id="L369"></a>func (a *compiler) compileType(b *block, typ ast.Expr) Type {
    <a id="L370"></a>tc := &amp;typeCompiler{a, b, noLateCheck};
    <a id="L371"></a>t := tc.compileType(typ, false);
    <a id="L372"></a>if !tc.lateCheck() {
        <a id="L373"></a>t = nil
    <a id="L374"></a>}
    <a id="L375"></a>return t;
<a id="L376"></a>}

<a id="L378"></a>func (a *compiler) compileTypeDecl(b *block, decl *ast.GenDecl) bool {
    <a id="L379"></a>ok := true;
    <a id="L380"></a>for _, spec := range decl.Specs {
        <a id="L381"></a>spec := spec.(*ast.TypeSpec);
        <a id="L382"></a><span class="comment">// Create incomplete type for this type</span>
        <a id="L383"></a>nt := b.DefineType(spec.Name.Value, spec.Name.Pos(), nil);
        <a id="L384"></a>if nt != nil {
            <a id="L385"></a>nt.(*NamedType).incomplete = true
        <a id="L386"></a>}
        <a id="L387"></a><span class="comment">// Compile type</span>
        <a id="L388"></a>tc := &amp;typeCompiler{a, b, noLateCheck};
        <a id="L389"></a>t := tc.compileType(spec.Type, false);
        <a id="L390"></a>if t == nil {
            <a id="L391"></a><span class="comment">// Create a placeholder type</span>
            <a id="L392"></a>ok = false
        <a id="L393"></a>}
        <a id="L394"></a><span class="comment">// Fill incomplete type</span>
        <a id="L395"></a>if nt != nil {
            <a id="L396"></a>nt.(*NamedType).Complete(t)
        <a id="L397"></a>}
        <a id="L398"></a><span class="comment">// Perform late type checking with complete type</span>
        <a id="L399"></a>if !tc.lateCheck() {
            <a id="L400"></a>ok = false;
            <a id="L401"></a>if nt != nil {
                <a id="L402"></a><span class="comment">// Make the type a placeholder</span>
                <a id="L403"></a>nt.(*NamedType).Def = nil
            <a id="L404"></a>}
        <a id="L405"></a>}
    <a id="L406"></a>}
    <a id="L407"></a>return ok;
<a id="L408"></a>}

<a id="L410"></a>func (a *compiler) compileFuncType(b *block, typ *ast.FuncType) *FuncDecl {
    <a id="L411"></a>tc := &amp;typeCompiler{a, b, noLateCheck};
    <a id="L412"></a>res := tc.compileFuncType(typ, false);
    <a id="L413"></a>if res != nil {
        <a id="L414"></a>if !tc.lateCheck() {
            <a id="L415"></a>res = nil
        <a id="L416"></a>}
    <a id="L417"></a>}
    <a id="L418"></a>return res;
<a id="L419"></a>}
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
