<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/ogle/rvalue.go</title>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/ogle/rvalue.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package ogle

<a id="L7"></a>import (
    <a id="L8"></a>&#34;debug/proc&#34;;
    <a id="L9"></a>&#34;exp/eval&#34;;
    <a id="L10"></a>&#34;fmt&#34;;
<a id="L11"></a>)

<a id="L13"></a><span class="comment">// A RemoteMismatchError occurs when an operation that requires two</span>
<a id="L14"></a><span class="comment">// identical remote processes is given different process.  For</span>
<a id="L15"></a><span class="comment">// example, this occurs when trying to set a pointer in one process to</span>
<a id="L16"></a><span class="comment">// point to something in another process.</span>
<a id="L17"></a>type RemoteMismatchError string

<a id="L19"></a>func (e RemoteMismatchError) String() string { return string(e) }

<a id="L21"></a><span class="comment">// A ReadOnlyError occurs when attempting to set or assign to a</span>
<a id="L22"></a><span class="comment">// read-only value.</span>
<a id="L23"></a>type ReadOnlyError string

<a id="L25"></a>func (e ReadOnlyError) String() string { return string(e) }

<a id="L27"></a><span class="comment">// A maker is a function that converts a remote address into an</span>
<a id="L28"></a><span class="comment">// interpreter Value.</span>
<a id="L29"></a>type maker func(remote) eval.Value

<a id="L31"></a>type remoteValue interface {
    <a id="L32"></a>addr() remote;
<a id="L33"></a>}

<a id="L35"></a><span class="comment">// remote represents an address in a remote process.</span>
<a id="L36"></a>type remote struct {
    <a id="L37"></a>base proc.Word;
    <a id="L38"></a>p    *Process;
<a id="L39"></a>}

<a id="L41"></a>func (v remote) Get(a aborter, size int) uint64 {
    <a id="L42"></a><span class="comment">// TODO(austin) This variable might temporarily be in a</span>
    <a id="L43"></a><span class="comment">// register.  We could trace the assembly back from the</span>
    <a id="L44"></a><span class="comment">// current PC, looking for the beginning of the function or a</span>
    <a id="L45"></a><span class="comment">// call (both of which guarantee that the variable is in</span>
    <a id="L46"></a><span class="comment">// memory), or an instruction that loads the variable into a</span>
    <a id="L47"></a><span class="comment">// register.</span>
    <a id="L48"></a><span class="comment">//</span>
    <a id="L49"></a><span class="comment">// TODO(austin) If this is a local variable, it might not be</span>
    <a id="L50"></a><span class="comment">// live at this PC.  In fact, because the compiler reuses</span>
    <a id="L51"></a><span class="comment">// slots, there might even be a different local variable at</span>
    <a id="L52"></a><span class="comment">// this location right now.  A simple solution to both</span>
    <a id="L53"></a><span class="comment">// problems is to include the range of PC&#39;s over which a local</span>
    <a id="L54"></a><span class="comment">// variable is live in the symbol table.</span>
    <a id="L55"></a><span class="comment">//</span>
    <a id="L56"></a><span class="comment">// TODO(austin) We need to prevent the remote garbage</span>
    <a id="L57"></a><span class="comment">// collector from collecting objects out from under us.</span>
    <a id="L58"></a>var arr [8]byte;
    <a id="L59"></a>buf := arr[0:size];
    <a id="L60"></a>_, err := v.p.Peek(v.base, buf);
    <a id="L61"></a>if err != nil {
        <a id="L62"></a>a.Abort(err)
    <a id="L63"></a>}
    <a id="L64"></a>return uint64(v.p.ToWord(buf));
<a id="L65"></a>}

<a id="L67"></a>func (v remote) Set(a aborter, size int, x uint64) {
    <a id="L68"></a>var arr [8]byte;
    <a id="L69"></a>buf := arr[0:size];
    <a id="L70"></a>v.p.FromWord(proc.Word(x), buf);
    <a id="L71"></a>_, err := v.p.Poke(v.base, buf);
    <a id="L72"></a>if err != nil {
        <a id="L73"></a>a.Abort(err)
    <a id="L74"></a>}
<a id="L75"></a>}

<a id="L77"></a>func (v remote) plus(x proc.Word) remote { return remote{v.base + x, v.p} }

<a id="L79"></a>func tryRVString(f func(a aborter) string) string {
    <a id="L80"></a>var s string;
    <a id="L81"></a>err := try(func(a aborter) { s = f(a) });
    <a id="L82"></a>if err != nil {
        <a id="L83"></a>return fmt.Sprintf(&#34;&lt;error: %v&gt;&#34;, err)
    <a id="L84"></a>}
    <a id="L85"></a>return s;
<a id="L86"></a>}

<a id="L88"></a><span class="comment">/*</span>
<a id="L89"></a><span class="comment"> * Bool</span>
<a id="L90"></a><span class="comment"> */</span>

<a id="L92"></a>type remoteBool struct {
    <a id="L93"></a>r remote;
<a id="L94"></a>}

<a id="L96"></a>func (v remoteBool) String() string {
    <a id="L97"></a>return tryRVString(func(a aborter) string { return fmt.Sprintf(&#34;%v&#34;, v.aGet(a)) })
<a id="L98"></a>}

<a id="L100"></a>func (v remoteBool) Assign(t *eval.Thread, o eval.Value) {
    <a id="L101"></a>v.Set(t, o.(eval.BoolValue).Get(t))
<a id="L102"></a>}

<a id="L104"></a>func (v remoteBool) Get(t *eval.Thread) bool { return v.aGet(t) }

<a id="L106"></a>func (v remoteBool) aGet(a aborter) bool { return v.r.Get(a, 1) != 0 }

<a id="L108"></a>func (v remoteBool) Set(t *eval.Thread, x bool) {
    <a id="L109"></a>v.aSet(t, x)
<a id="L110"></a>}

<a id="L112"></a>func (v remoteBool) aSet(a aborter, x bool) {
    <a id="L113"></a>if x {
        <a id="L114"></a>v.r.Set(a, 1, 1)
    <a id="L115"></a>} else {
        <a id="L116"></a>v.r.Set(a, 1, 0)
    <a id="L117"></a>}
<a id="L118"></a>}

<a id="L120"></a>func (v remoteBool) addr() remote { return v.r }

<a id="L122"></a>func mkBool(r remote) eval.Value { return remoteBool{r} }

<a id="L124"></a><span class="comment">/*</span>
<a id="L125"></a><span class="comment"> * Uint</span>
<a id="L126"></a><span class="comment"> */</span>

<a id="L128"></a>type remoteUint struct {
    <a id="L129"></a>r    remote;
    <a id="L130"></a>size int;
<a id="L131"></a>}

<a id="L133"></a>func (v remoteUint) String() string {
    <a id="L134"></a>return tryRVString(func(a aborter) string { return fmt.Sprintf(&#34;%v&#34;, v.aGet(a)) })
<a id="L135"></a>}

<a id="L137"></a>func (v remoteUint) Assign(t *eval.Thread, o eval.Value) {
    <a id="L138"></a>v.Set(t, o.(eval.UintValue).Get(t))
<a id="L139"></a>}

<a id="L141"></a>func (v remoteUint) Get(t *eval.Thread) uint64 {
    <a id="L142"></a>return v.aGet(t)
<a id="L143"></a>}

<a id="L145"></a>func (v remoteUint) aGet(a aborter) uint64 { return v.r.Get(a, v.size) }

<a id="L147"></a>func (v remoteUint) Set(t *eval.Thread, x uint64) {
    <a id="L148"></a>v.aSet(t, x)
<a id="L149"></a>}

<a id="L151"></a>func (v remoteUint) aSet(a aborter, x uint64) { v.r.Set(a, v.size, x) }

<a id="L153"></a>func (v remoteUint) addr() remote { return v.r }

<a id="L155"></a>func mkUint8(r remote) eval.Value { return remoteUint{r, 1} }

<a id="L157"></a>func mkUint16(r remote) eval.Value { return remoteUint{r, 2} }

<a id="L159"></a>func mkUint32(r remote) eval.Value { return remoteUint{r, 4} }

<a id="L161"></a>func mkUint64(r remote) eval.Value { return remoteUint{r, 8} }

<a id="L163"></a>func mkUint(r remote) eval.Value { return remoteUint{r, r.p.IntSize()} }

<a id="L165"></a>func mkUintptr(r remote) eval.Value { return remoteUint{r, r.p.PtrSize()} }

<a id="L167"></a><span class="comment">/*</span>
<a id="L168"></a><span class="comment"> * Int</span>
<a id="L169"></a><span class="comment"> */</span>

<a id="L171"></a>type remoteInt struct {
    <a id="L172"></a>r    remote;
    <a id="L173"></a>size int;
<a id="L174"></a>}

<a id="L176"></a>func (v remoteInt) String() string {
    <a id="L177"></a>return tryRVString(func(a aborter) string { return fmt.Sprintf(&#34;%v&#34;, v.aGet(a)) })
<a id="L178"></a>}

<a id="L180"></a>func (v remoteInt) Assign(t *eval.Thread, o eval.Value) {
    <a id="L181"></a>v.Set(t, o.(eval.IntValue).Get(t))
<a id="L182"></a>}

<a id="L184"></a>func (v remoteInt) Get(t *eval.Thread) int64 { return v.aGet(t) }

<a id="L186"></a>func (v remoteInt) aGet(a aborter) int64 { return int64(v.r.Get(a, v.size)) }

<a id="L188"></a>func (v remoteInt) Set(t *eval.Thread, x int64) {
    <a id="L189"></a>v.aSet(t, x)
<a id="L190"></a>}

<a id="L192"></a>func (v remoteInt) aSet(a aborter, x int64) { v.r.Set(a, v.size, uint64(x)) }

<a id="L194"></a>func (v remoteInt) addr() remote { return v.r }

<a id="L196"></a>func mkInt8(r remote) eval.Value { return remoteInt{r, 1} }

<a id="L198"></a>func mkInt16(r remote) eval.Value { return remoteInt{r, 2} }

<a id="L200"></a>func mkInt32(r remote) eval.Value { return remoteInt{r, 4} }

<a id="L202"></a>func mkInt64(r remote) eval.Value { return remoteInt{r, 8} }

<a id="L204"></a>func mkInt(r remote) eval.Value { return remoteInt{r, r.p.IntSize()} }

<a id="L206"></a><span class="comment">/*</span>
<a id="L207"></a><span class="comment"> * Float</span>
<a id="L208"></a><span class="comment"> */</span>

<a id="L210"></a>type remoteFloat struct {
    <a id="L211"></a>r    remote;
    <a id="L212"></a>size int;
<a id="L213"></a>}

<a id="L215"></a>func (v remoteFloat) String() string {
    <a id="L216"></a>return tryRVString(func(a aborter) string { return fmt.Sprintf(&#34;%v&#34;, v.aGet(a)) })
<a id="L217"></a>}

<a id="L219"></a>func (v remoteFloat) Assign(t *eval.Thread, o eval.Value) {
    <a id="L220"></a>v.Set(t, o.(eval.FloatValue).Get(t))
<a id="L221"></a>}

<a id="L223"></a>func (v remoteFloat) Get(t *eval.Thread) float64 {
    <a id="L224"></a>return v.aGet(t)
<a id="L225"></a>}

<a id="L227"></a>func (v remoteFloat) aGet(a aborter) float64 {
    <a id="L228"></a>bits := v.r.Get(a, v.size);
    <a id="L229"></a>switch v.size {
    <a id="L230"></a>case 4:
        <a id="L231"></a>return float64(v.r.p.ToFloat32(uint32(bits)))
    <a id="L232"></a>case 8:
        <a id="L233"></a>return v.r.p.ToFloat64(bits)
    <a id="L234"></a>}
    <a id="L235"></a>panic(&#34;Unexpected float size &#34;, v.size);
<a id="L236"></a>}

<a id="L238"></a>func (v remoteFloat) Set(t *eval.Thread, x float64) {
    <a id="L239"></a>v.aSet(t, x)
<a id="L240"></a>}

<a id="L242"></a>func (v remoteFloat) aSet(a aborter, x float64) {
    <a id="L243"></a>var bits uint64;
    <a id="L244"></a>switch v.size {
    <a id="L245"></a>case 4:
        <a id="L246"></a>bits = uint64(v.r.p.FromFloat32(float32(x)))
    <a id="L247"></a>case 8:
        <a id="L248"></a>bits = v.r.p.FromFloat64(x)
    <a id="L249"></a>default:
        <a id="L250"></a>panic(&#34;Unexpected float size &#34;, v.size)
    <a id="L251"></a>}
    <a id="L252"></a>v.r.Set(a, v.size, bits);
<a id="L253"></a>}

<a id="L255"></a>func (v remoteFloat) addr() remote { return v.r }

<a id="L257"></a>func mkFloat32(r remote) eval.Value { return remoteFloat{r, 4} }

<a id="L259"></a>func mkFloat64(r remote) eval.Value { return remoteFloat{r, 8} }

<a id="L261"></a>func mkFloat(r remote) eval.Value { return remoteFloat{r, r.p.FloatSize()} }

<a id="L263"></a><span class="comment">/*</span>
<a id="L264"></a><span class="comment"> * String</span>
<a id="L265"></a><span class="comment"> */</span>

<a id="L267"></a>type remoteString struct {
    <a id="L268"></a>r remote;
<a id="L269"></a>}

<a id="L271"></a>func (v remoteString) String() string {
    <a id="L272"></a>return tryRVString(func(a aborter) string { return v.aGet(a) })
<a id="L273"></a>}

<a id="L275"></a>func (v remoteString) Assign(t *eval.Thread, o eval.Value) {
    <a id="L276"></a>v.Set(t, o.(eval.StringValue).Get(t))
<a id="L277"></a>}

<a id="L279"></a>func (v remoteString) Get(t *eval.Thread) string {
    <a id="L280"></a>return v.aGet(t)
<a id="L281"></a>}

<a id="L283"></a>func (v remoteString) aGet(a aborter) string {
    <a id="L284"></a>rs := v.r.p.runtime.String.mk(v.r).(remoteStruct);
    <a id="L285"></a>str := proc.Word(rs.field(v.r.p.f.String.Str).(remoteUint).aGet(a));
    <a id="L286"></a>len := rs.field(v.r.p.f.String.Len).(remoteInt).aGet(a);

    <a id="L288"></a>bytes := make([]uint8, len);
    <a id="L289"></a>_, err := v.r.p.Peek(str, bytes);
    <a id="L290"></a>if err != nil {
        <a id="L291"></a>a.Abort(err)
    <a id="L292"></a>}
    <a id="L293"></a>return string(bytes);
<a id="L294"></a>}

<a id="L296"></a>func (v remoteString) Set(t *eval.Thread, x string) {
    <a id="L297"></a>v.aSet(t, x)
<a id="L298"></a>}

<a id="L300"></a>func (v remoteString) aSet(a aborter, x string) {
    <a id="L301"></a><span class="comment">// TODO(austin) This isn&#39;t generally possible without the</span>
    <a id="L302"></a><span class="comment">// ability to allocate remote memory.</span>
    <a id="L303"></a>a.Abort(ReadOnlyError(&#34;remote strings cannot be assigned to&#34;))
<a id="L304"></a>}

<a id="L306"></a>func mkString(r remote) eval.Value { return remoteString{r} }

<a id="L308"></a><span class="comment">/*</span>
<a id="L309"></a><span class="comment"> * Array</span>
<a id="L310"></a><span class="comment"> */</span>

<a id="L312"></a>type remoteArray struct {
    <a id="L313"></a>r        remote;
    <a id="L314"></a>len      int64;
    <a id="L315"></a>elemType *remoteType;
<a id="L316"></a>}

<a id="L318"></a>func (v remoteArray) String() string {
    <a id="L319"></a>res := &#34;{&#34;;
    <a id="L320"></a>for i := int64(0); i &lt; v.len; i++ {
        <a id="L321"></a>if i &gt; 0 {
            <a id="L322"></a>res += &#34;, &#34;
        <a id="L323"></a>}
        <a id="L324"></a>res += v.elem(i).String();
    <a id="L325"></a>}
    <a id="L326"></a>return res + &#34;}&#34;;
<a id="L327"></a>}

<a id="L329"></a>func (v remoteArray) Assign(t *eval.Thread, o eval.Value) {
    <a id="L330"></a><span class="comment">// TODO(austin) Could do a bigger memcpy if o is a</span>
    <a id="L331"></a><span class="comment">// remoteArray in the same Process.</span>
    <a id="L332"></a>oa := o.(eval.ArrayValue);
    <a id="L333"></a>for i := int64(0); i &lt; v.len; i++ {
        <a id="L334"></a>v.Elem(t, i).Assign(t, oa.Elem(t, i))
    <a id="L335"></a>}
<a id="L336"></a>}

<a id="L338"></a>func (v remoteArray) Get(t *eval.Thread) eval.ArrayValue {
    <a id="L339"></a>return v
<a id="L340"></a>}

<a id="L342"></a>func (v remoteArray) Elem(t *eval.Thread, i int64) eval.Value {
    <a id="L343"></a>return v.elem(i)
<a id="L344"></a>}

<a id="L346"></a>func (v remoteArray) elem(i int64) eval.Value {
    <a id="L347"></a>return v.elemType.mk(v.r.plus(proc.Word(int64(v.elemType.size) * i)))
<a id="L348"></a>}

<a id="L350"></a>func (v remoteArray) Sub(i int64, len int64) eval.ArrayValue {
    <a id="L351"></a>return remoteArray{v.r.plus(proc.Word(int64(v.elemType.size) * i)), len, v.elemType}
<a id="L352"></a>}

<a id="L354"></a><span class="comment">/*</span>
<a id="L355"></a><span class="comment"> * Struct</span>
<a id="L356"></a><span class="comment"> */</span>

<a id="L358"></a>type remoteStruct struct {
    <a id="L359"></a>r      remote;
    <a id="L360"></a>layout []remoteStructField;
<a id="L361"></a>}

<a id="L363"></a>type remoteStructField struct {
    <a id="L364"></a>offset    int;
    <a id="L365"></a>fieldType *remoteType;
<a id="L366"></a>}

<a id="L368"></a>func (v remoteStruct) String() string {
    <a id="L369"></a>res := &#34;{&#34;;
    <a id="L370"></a>for i := range v.layout {
        <a id="L371"></a>if i &gt; 0 {
            <a id="L372"></a>res += &#34;, &#34;
        <a id="L373"></a>}
        <a id="L374"></a>res += v.field(i).String();
    <a id="L375"></a>}
    <a id="L376"></a>return res + &#34;}&#34;;
<a id="L377"></a>}

<a id="L379"></a>func (v remoteStruct) Assign(t *eval.Thread, o eval.Value) {
    <a id="L380"></a><span class="comment">// TODO(austin) Could do a bigger memcpy.</span>
    <a id="L381"></a>oa := o.(eval.StructValue);
    <a id="L382"></a>l := len(v.layout);
    <a id="L383"></a>for i := 0; i &lt; l; i++ {
        <a id="L384"></a>v.Field(t, i).Assign(t, oa.Field(t, i))
    <a id="L385"></a>}
<a id="L386"></a>}

<a id="L388"></a>func (v remoteStruct) Get(t *eval.Thread) eval.StructValue {
    <a id="L389"></a>return v
<a id="L390"></a>}

<a id="L392"></a>func (v remoteStruct) Field(t *eval.Thread, i int) eval.Value {
    <a id="L393"></a>return v.field(i)
<a id="L394"></a>}

<a id="L396"></a>func (v remoteStruct) field(i int) eval.Value {
    <a id="L397"></a>f := &amp;v.layout[i];
    <a id="L398"></a>return f.fieldType.mk(v.r.plus(proc.Word(f.offset)));
<a id="L399"></a>}

<a id="L401"></a>func (v remoteStruct) addr() remote { return v.r }

<a id="L403"></a><span class="comment">/*</span>
<a id="L404"></a><span class="comment"> * Pointer</span>
<a id="L405"></a><span class="comment"> */</span>

<a id="L407"></a><span class="comment">// TODO(austin) Comparing two remote pointers for equality in the</span>
<a id="L408"></a><span class="comment">// interpreter will crash it because the Value&#39;s returned from</span>
<a id="L409"></a><span class="comment">// remotePtr.Get() will be structs.</span>

<a id="L411"></a>type remotePtr struct {
    <a id="L412"></a>r        remote;
    <a id="L413"></a>elemType *remoteType;
<a id="L414"></a>}

<a id="L416"></a>func (v remotePtr) String() string {
    <a id="L417"></a>return tryRVString(func(a aborter) string {
        <a id="L418"></a>e := v.aGet(a);
        <a id="L419"></a>if e == nil {
            <a id="L420"></a>return &#34;&lt;nil&gt;&#34;
        <a id="L421"></a>}
        <a id="L422"></a>return &#34;&amp;&#34; + e.String();
    <a id="L423"></a>})
<a id="L424"></a>}

<a id="L426"></a>func (v remotePtr) Assign(t *eval.Thread, o eval.Value) {
    <a id="L427"></a>v.Set(t, o.(eval.PtrValue).Get(t))
<a id="L428"></a>}

<a id="L430"></a>func (v remotePtr) Get(t *eval.Thread) eval.Value {
    <a id="L431"></a>return v.aGet(t)
<a id="L432"></a>}

<a id="L434"></a>func (v remotePtr) aGet(a aborter) eval.Value {
    <a id="L435"></a>addr := proc.Word(v.r.Get(a, v.r.p.PtrSize()));
    <a id="L436"></a>if addr == 0 {
        <a id="L437"></a>return nil
    <a id="L438"></a>}
    <a id="L439"></a>return v.elemType.mk(remote{addr, v.r.p});
<a id="L440"></a>}

<a id="L442"></a>func (v remotePtr) Set(t *eval.Thread, x eval.Value) {
    <a id="L443"></a>v.aSet(t, x)
<a id="L444"></a>}

<a id="L446"></a>func (v remotePtr) aSet(a aborter, x eval.Value) {
    <a id="L447"></a>if x == nil {
        <a id="L448"></a>v.r.Set(a, v.r.p.PtrSize(), 0);
        <a id="L449"></a>return;
    <a id="L450"></a>}
    <a id="L451"></a>xr, ok := x.(remoteValue);
    <a id="L452"></a>if !ok || v.r.p != xr.addr().p {
        <a id="L453"></a>a.Abort(RemoteMismatchError(&#34;remote pointer must point within the same process&#34;))
    <a id="L454"></a>}
    <a id="L455"></a>v.r.Set(a, v.r.p.PtrSize(), uint64(xr.addr().base));
<a id="L456"></a>}

<a id="L458"></a>func (v remotePtr) addr() remote { return v.r }

<a id="L460"></a><span class="comment">/*</span>
<a id="L461"></a><span class="comment"> * Slice</span>
<a id="L462"></a><span class="comment"> */</span>

<a id="L464"></a>type remoteSlice struct {
    <a id="L465"></a>r        remote;
    <a id="L466"></a>elemType *remoteType;
<a id="L467"></a>}

<a id="L469"></a>func (v remoteSlice) String() string {
    <a id="L470"></a>return tryRVString(func(a aborter) string {
        <a id="L471"></a>b := v.aGet(a).Base;
        <a id="L472"></a>if b == nil {
            <a id="L473"></a>return &#34;&lt;nil&gt;&#34;
        <a id="L474"></a>}
        <a id="L475"></a>return b.String();
    <a id="L476"></a>})
<a id="L477"></a>}

<a id="L479"></a>func (v remoteSlice) Assign(t *eval.Thread, o eval.Value) {
    <a id="L480"></a>v.Set(t, o.(eval.SliceValue).Get(t))
<a id="L481"></a>}

<a id="L483"></a>func (v remoteSlice) Get(t *eval.Thread) eval.Slice {
    <a id="L484"></a>return v.aGet(t)
<a id="L485"></a>}

<a id="L487"></a>func (v remoteSlice) aGet(a aborter) eval.Slice {
    <a id="L488"></a>rs := v.r.p.runtime.Slice.mk(v.r).(remoteStruct);
    <a id="L489"></a>base := proc.Word(rs.field(v.r.p.f.Slice.Array).(remoteUint).aGet(a));
    <a id="L490"></a>nel := rs.field(v.r.p.f.Slice.Len).(remoteInt).aGet(a);
    <a id="L491"></a>cap := rs.field(v.r.p.f.Slice.Cap).(remoteInt).aGet(a);
    <a id="L492"></a>if base == 0 {
        <a id="L493"></a>return eval.Slice{nil, nel, cap}
    <a id="L494"></a>}
    <a id="L495"></a>return eval.Slice{remoteArray{remote{base, v.r.p}, nel, v.elemType}, nel, cap};
<a id="L496"></a>}

<a id="L498"></a>func (v remoteSlice) Set(t *eval.Thread, x eval.Slice) {
    <a id="L499"></a>v.aSet(t, x)
<a id="L500"></a>}

<a id="L502"></a>func (v remoteSlice) aSet(a aborter, x eval.Slice) {
    <a id="L503"></a>rs := v.r.p.runtime.Slice.mk(v.r).(remoteStruct);
    <a id="L504"></a>if x.Base == nil {
        <a id="L505"></a>rs.field(v.r.p.f.Slice.Array).(remoteUint).aSet(a, 0)
    <a id="L506"></a>} else {
        <a id="L507"></a>ar, ok := x.Base.(remoteArray);
        <a id="L508"></a>if !ok || v.r.p != ar.r.p {
            <a id="L509"></a>a.Abort(RemoteMismatchError(&#34;remote slice must point within the same process&#34;))
        <a id="L510"></a>}
        <a id="L511"></a>rs.field(v.r.p.f.Slice.Array).(remoteUint).aSet(a, uint64(ar.r.base));
    <a id="L512"></a>}
    <a id="L513"></a>rs.field(v.r.p.f.Slice.Len).(remoteInt).aSet(a, x.Len);
    <a id="L514"></a>rs.field(v.r.p.f.Slice.Cap).(remoteInt).aSet(a, x.Cap);
<a id="L515"></a>}
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
