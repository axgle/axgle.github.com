<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/ogle/vars.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/exp/ogle/vars.go</h1>

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
    <a id="L8"></a>&#34;debug/gosym&#34;;
    <a id="L9"></a>&#34;debug/proc&#34;;
    <a id="L10"></a>&#34;exp/eval&#34;;
    <a id="L11"></a>&#34;log&#34;;
    <a id="L12"></a>&#34;os&#34;;
<a id="L13"></a>)

<a id="L15"></a><span class="comment">/*</span>
<a id="L16"></a><span class="comment"> * Remote frame pointers</span>
<a id="L17"></a><span class="comment"> */</span>

<a id="L19"></a><span class="comment">// A NotOnStack error occurs when attempting to access a variable in a</span>
<a id="L20"></a><span class="comment">// remote frame where that remote frame is not on the current stack.</span>
<a id="L21"></a>type NotOnStack struct {
    <a id="L22"></a>Fn        *gosym.Func;
    <a id="L23"></a>Goroutine *Goroutine;
<a id="L24"></a>}

<a id="L26"></a>func (e NotOnStack) String() string {
    <a id="L27"></a>return &#34;function &#34; + e.Fn.Name + &#34; not on &#34; + e.Goroutine.String() + &#34;&#39;s stack&#34;
<a id="L28"></a>}

<a id="L30"></a><span class="comment">// A remoteFramePtr is an implementation of eval.PtrValue that</span>
<a id="L31"></a><span class="comment">// represents a pointer to a function frame in a remote process.  When</span>
<a id="L32"></a><span class="comment">// accessed, this locates the function on the current goroutine&#39;s</span>
<a id="L33"></a><span class="comment">// stack and returns a structure containing the local variables of</span>
<a id="L34"></a><span class="comment">// that function.</span>
<a id="L35"></a>type remoteFramePtr struct {
    <a id="L36"></a>p   *Process;
    <a id="L37"></a>fn  *gosym.Func;
    <a id="L38"></a>rt  *remoteType;
<a id="L39"></a>}

<a id="L41"></a>func (v remoteFramePtr) String() string {
    <a id="L42"></a><span class="comment">// TODO(austin): This could be a really awesome string method</span>
    <a id="L43"></a>return &#34;&lt;remote frame&gt;&#34;
<a id="L44"></a>}

<a id="L46"></a>func (v remoteFramePtr) Assign(t *eval.Thread, o eval.Value) {
    <a id="L47"></a>v.Set(t, o.(eval.PtrValue).Get(t))
<a id="L48"></a>}

<a id="L50"></a>func (v remoteFramePtr) Get(t *eval.Thread) eval.Value {
    <a id="L51"></a>g := v.p.curGoroutine;
    <a id="L52"></a>if g == nil || g.frame == nil {
        <a id="L53"></a>t.Abort(NoCurrentGoroutine{})
    <a id="L54"></a>}

    <a id="L56"></a>for f := g.frame; f != nil; f = f.aOuter(t) {
        <a id="L57"></a>if f.fn != v.fn {
            <a id="L58"></a>continue
        <a id="L59"></a>}

        <a id="L61"></a><span class="comment">// TODO(austin): Register for shootdown with f</span>
        <a id="L62"></a>return v.rt.mk(remote{f.fp, v.p});
    <a id="L63"></a>}

    <a id="L65"></a>t.Abort(NotOnStack{v.fn, g});
    <a id="L66"></a>panic();
<a id="L67"></a>}

<a id="L69"></a>func (v remoteFramePtr) Set(t *eval.Thread, x eval.Value) {
    <a id="L70"></a><span class="comment">// Theoretically this could be a static error.  If remote</span>
    <a id="L71"></a><span class="comment">// packages were packages, remote frames could just be defined</span>
    <a id="L72"></a><span class="comment">// as constants.</span>
    <a id="L73"></a>t.Abort(ReadOnlyError(&#34;remote frames cannot be assigned to&#34;))
<a id="L74"></a>}

<a id="L76"></a><span class="comment">/*</span>
<a id="L77"></a><span class="comment"> * Remote packages</span>
<a id="L78"></a><span class="comment"> */</span>

<a id="L80"></a><span class="comment">// TODO(austin): Remote packages are implemented as structs right now,</span>
<a id="L81"></a><span class="comment">// which has some weird consequences.  You can attempt to assign to a</span>
<a id="L82"></a><span class="comment">// remote package.  It also produces terrible error messages.</span>
<a id="L83"></a><span class="comment">// Ideally, these would actually be packages, but somehow first-class</span>
<a id="L84"></a><span class="comment">// so they could be assigned to other names.</span>

<a id="L86"></a><span class="comment">// A remotePackage is an implementation of eval.StructValue that</span>
<a id="L87"></a><span class="comment">// represents a package in a remote process.  It&#39;s essentially a</span>
<a id="L88"></a><span class="comment">// regular struct, except it cannot be assigned to.</span>
<a id="L89"></a>type remotePackage struct {
    <a id="L90"></a>defs []eval.Value;
<a id="L91"></a>}

<a id="L93"></a>func (v remotePackage) String() string { return &#34;&lt;remote package&gt;&#34; }

<a id="L95"></a>func (v remotePackage) Assign(t *eval.Thread, o eval.Value) {
    <a id="L96"></a>t.Abort(ReadOnlyError(&#34;remote packages cannot be assigned to&#34;))
<a id="L97"></a>}

<a id="L99"></a>func (v remotePackage) Get(t *eval.Thread) eval.StructValue {
    <a id="L100"></a>return v
<a id="L101"></a>}

<a id="L103"></a>func (v remotePackage) Field(t *eval.Thread, i int) eval.Value {
    <a id="L104"></a>return v.defs[i]
<a id="L105"></a>}

<a id="L107"></a><span class="comment">/*</span>
<a id="L108"></a><span class="comment"> * Remote variables</span>
<a id="L109"></a><span class="comment"> */</span>

<a id="L111"></a><span class="comment">// populateWorld defines constants in the given world for each package</span>
<a id="L112"></a><span class="comment">// in this process.  These packages are structs that, in turn, contain</span>
<a id="L113"></a><span class="comment">// fields for each global and function in that package.</span>
<a id="L114"></a>func (p *Process) populateWorld(w *eval.World) os.Error {
    <a id="L115"></a>type def struct {
        <a id="L116"></a>t   eval.Type;
        <a id="L117"></a>v   eval.Value;
    <a id="L118"></a>}
    <a id="L119"></a>packages := make(map[string]map[string]def);

    <a id="L121"></a>for _, s := range p.syms.Syms {
        <a id="L122"></a>if s.ReceiverName() != &#34;&#34; {
            <a id="L123"></a><span class="comment">// TODO(austin)</span>
            <a id="L124"></a>continue
        <a id="L125"></a>}

        <a id="L127"></a><span class="comment">// Package</span>
        <a id="L128"></a>pkgName := s.PackageName();
        <a id="L129"></a>switch pkgName {
        <a id="L130"></a>case &#34;&#34;, &#34;type&#34;, &#34;extratype&#34;, &#34;string&#34;, &#34;go&#34;:
            <a id="L131"></a><span class="comment">// &#34;go&#34; is really &#34;go.string&#34;</span>
            <a id="L132"></a>continue
        <a id="L133"></a>}
        <a id="L134"></a>pkg, ok := packages[pkgName];
        <a id="L135"></a>if !ok {
            <a id="L136"></a>pkg = make(map[string]def);
            <a id="L137"></a>packages[pkgName] = pkg;
        <a id="L138"></a>}

        <a id="L140"></a><span class="comment">// Symbol name</span>
        <a id="L141"></a>name := s.BaseName();
        <a id="L142"></a>if _, ok := pkg[name]; ok {
            <a id="L143"></a>log.Stderrf(&#34;Multiple definitions of symbol %s&#34;, s.Name);
            <a id="L144"></a>continue;
        <a id="L145"></a>}

        <a id="L147"></a><span class="comment">// Symbol type</span>
        <a id="L148"></a>rt, err := p.typeOfSym(&amp;s);
        <a id="L149"></a>if err != nil {
            <a id="L150"></a>return err
        <a id="L151"></a>}

        <a id="L153"></a><span class="comment">// Definition</span>
        <a id="L154"></a>switch s.Type {
        <a id="L155"></a>case &#39;D&#39;, &#39;d&#39;, &#39;B&#39;, &#39;b&#39;:
            <a id="L156"></a><span class="comment">// Global variable</span>
            <a id="L157"></a>if rt == nil {
                <a id="L158"></a>continue
            <a id="L159"></a>}
            <a id="L160"></a>pkg[name] = def{rt.Type, rt.mk(remote{proc.Word(s.Value), p})};

        <a id="L162"></a>case &#39;T&#39;, &#39;t&#39;, &#39;L&#39;, &#39;l&#39;:
            <a id="L163"></a><span class="comment">// Function</span>
            <a id="L164"></a>s := s.Func;
            <a id="L165"></a><span class="comment">// TODO(austin): Ideally, this would *also* be</span>
            <a id="L166"></a><span class="comment">// callable.  How does that interact with type</span>
            <a id="L167"></a><span class="comment">// conversion syntax?</span>
            <a id="L168"></a>rt, err := p.makeFrameType(s);
            <a id="L169"></a>if err != nil {
                <a id="L170"></a>return err
            <a id="L171"></a>}
            <a id="L172"></a>pkg[name] = def{eval.NewPtrType(rt.Type), remoteFramePtr{p, s, rt}};
        <a id="L173"></a>}
    <a id="L174"></a>}

    <a id="L176"></a><span class="comment">// TODO(austin): Define remote types</span>

    <a id="L178"></a><span class="comment">// Define packages</span>
    <a id="L179"></a>for pkgName, defs := range packages {
        <a id="L180"></a>fields := make([]eval.StructField, len(defs));
        <a id="L181"></a>vals := make([]eval.Value, len(defs));
        <a id="L182"></a>i := 0;
        <a id="L183"></a>for name, def := range defs {
            <a id="L184"></a>fields[i].Name = name;
            <a id="L185"></a>fields[i].Type = def.t;
            <a id="L186"></a>vals[i] = def.v;
            <a id="L187"></a>i++;
        <a id="L188"></a>}
        <a id="L189"></a>pkgType := eval.NewStructType(fields);
        <a id="L190"></a>pkgVal := remotePackage{vals};

        <a id="L192"></a>err := w.DefineConst(pkgName, pkgType, pkgVal);
        <a id="L193"></a>if err != nil {
            <a id="L194"></a>log.Stderrf(&#34;while defining package %s: %v&#34;, pkgName, err)
        <a id="L195"></a>}
    <a id="L196"></a>}

    <a id="L198"></a>return nil;
<a id="L199"></a>}

<a id="L201"></a><span class="comment">// typeOfSym returns the type associated with a symbol.  If the symbol</span>
<a id="L202"></a><span class="comment">// has no type, returns nil.</span>
<a id="L203"></a>func (p *Process) typeOfSym(s *gosym.Sym) (*remoteType, os.Error) {
    <a id="L204"></a>if s.GoType == 0 {
        <a id="L205"></a>return nil, nil
    <a id="L206"></a>}
    <a id="L207"></a>addr := proc.Word(s.GoType);
    <a id="L208"></a>var rt *remoteType;
    <a id="L209"></a>err := try(func(a aborter) { rt = parseRemoteType(a, p.runtime.Type.mk(remote{addr, p}).(remoteStruct)) });
    <a id="L210"></a>if err != nil {
        <a id="L211"></a>return nil, err
    <a id="L212"></a>}
    <a id="L213"></a>return rt, nil;
<a id="L214"></a>}

<a id="L216"></a><span class="comment">// makeFrameType constructs a struct type for the frame of a function.</span>
<a id="L217"></a><span class="comment">// The offsets in this struct type are such that the struct can be</span>
<a id="L218"></a><span class="comment">// instantiated at this function&#39;s frame pointer.</span>
<a id="L219"></a>func (p *Process) makeFrameType(s *gosym.Func) (*remoteType, os.Error) {
    <a id="L220"></a>n := len(s.Params) + len(s.Locals);
    <a id="L221"></a>fields := make([]eval.StructField, n);
    <a id="L222"></a>layout := make([]remoteStructField, n);
    <a id="L223"></a>i := 0;

    <a id="L225"></a><span class="comment">// TODO(austin): There can be multiple locals/parameters with</span>
    <a id="L226"></a><span class="comment">// the same name.  We probably need liveness information to do</span>
    <a id="L227"></a><span class="comment">// anything about this.  Once we have that, perhaps we give</span>
    <a id="L228"></a><span class="comment">// such fields interface{} type?  Or perhaps we disambiguate</span>
    <a id="L229"></a><span class="comment">// the names with numbers.  Disambiguation is annoying for</span>
    <a id="L230"></a><span class="comment">// things like &#34;i&#34;, where there&#39;s an obvious right answer.</span>

    <a id="L232"></a>for _, param := range s.Params {
        <a id="L233"></a>rt, err := p.typeOfSym(param);
        <a id="L234"></a>if err != nil {
            <a id="L235"></a>return nil, err
        <a id="L236"></a>}
        <a id="L237"></a>if rt == nil {
            <a id="L238"></a><span class="comment">//fmt.Printf(&#34; (no type)\n&#34;);</span>
            <a id="L239"></a>continue
        <a id="L240"></a>}
        <a id="L241"></a><span class="comment">// TODO(austin): Why do local variables carry their</span>
        <a id="L242"></a><span class="comment">// package name?</span>
        <a id="L243"></a>fields[i].Name = param.BaseName();
        <a id="L244"></a>fields[i].Type = rt.Type;
        <a id="L245"></a><span class="comment">// Parameters have positive offsets from FP</span>
        <a id="L246"></a>layout[i].offset = int(param.Value);
        <a id="L247"></a>layout[i].fieldType = rt;
        <a id="L248"></a>i++;
    <a id="L249"></a>}

    <a id="L251"></a>for _, local := range s.Locals {
        <a id="L252"></a>rt, err := p.typeOfSym(local);
        <a id="L253"></a>if err != nil {
            <a id="L254"></a>return nil, err
        <a id="L255"></a>}
        <a id="L256"></a>if rt == nil {
            <a id="L257"></a>continue
        <a id="L258"></a>}
        <a id="L259"></a>fields[i].Name = local.BaseName();
        <a id="L260"></a>fields[i].Type = rt.Type;
        <a id="L261"></a><span class="comment">// Locals have negative offsets from FP - PtrSize</span>
        <a id="L262"></a>layout[i].offset = -int(local.Value) - p.PtrSize();
        <a id="L263"></a>layout[i].fieldType = rt;
        <a id="L264"></a>i++;
    <a id="L265"></a>}

    <a id="L267"></a>fields = fields[0:i];
    <a id="L268"></a>layout = layout[0:i];
    <a id="L269"></a>t := eval.NewStructType(fields);
    <a id="L270"></a>mk := func(r remote) eval.Value { return remoteStruct{r, layout} };
    <a id="L271"></a>return &amp;remoteType{t, 0, 0, mk}, nil;
<a id="L272"></a>}
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
