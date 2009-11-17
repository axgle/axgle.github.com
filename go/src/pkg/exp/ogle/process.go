<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/ogle/process.go</title>

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
	<li>Thu Nov 12 15:47:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/ogle/process.go</h1>

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
    <a id="L8"></a>&#34;debug/elf&#34;;
    <a id="L9"></a>&#34;debug/gosym&#34;;
    <a id="L10"></a>&#34;debug/proc&#34;;
    <a id="L11"></a>&#34;exp/eval&#34;;
    <a id="L12"></a>&#34;fmt&#34;;
    <a id="L13"></a>&#34;log&#34;;
    <a id="L14"></a>&#34;os&#34;;
    <a id="L15"></a>&#34;reflect&#34;;
<a id="L16"></a>)

<a id="L18"></a><span class="comment">// A FormatError indicates a failure to process information in or</span>
<a id="L19"></a><span class="comment">// about a remote process, such as unexpected or missing information</span>
<a id="L20"></a><span class="comment">// in the object file or runtime structures.</span>
<a id="L21"></a>type FormatError string

<a id="L23"></a>func (e FormatError) String() string { return string(e) }

<a id="L25"></a><span class="comment">// An UnknownArchitecture occurs when trying to load an object file</span>
<a id="L26"></a><span class="comment">// that indicates an architecture not supported by the debugger.</span>
<a id="L27"></a>type UnknownArchitecture elf.Machine

<a id="L29"></a>func (e UnknownArchitecture) String() string {
    <a id="L30"></a>return &#34;unknown architecture: &#34; + elf.Machine(e).String()
<a id="L31"></a>}

<a id="L33"></a><span class="comment">// A ProcessNotStopped error occurs when attempting to read or write</span>
<a id="L34"></a><span class="comment">// memory or registers of a process that is not stopped.</span>
<a id="L35"></a>type ProcessNotStopped struct{}

<a id="L37"></a>func (e ProcessNotStopped) String() string { return &#34;process not stopped&#34; }

<a id="L39"></a><span class="comment">// An UnknownGoroutine error is an internal error representing an</span>
<a id="L40"></a><span class="comment">// unrecognized G structure pointer.</span>
<a id="L41"></a>type UnknownGoroutine struct {
    <a id="L42"></a>OSThread  proc.Thread;
    <a id="L43"></a>Goroutine proc.Word;
<a id="L44"></a>}

<a id="L46"></a>func (e UnknownGoroutine) String() string {
    <a id="L47"></a>return fmt.Sprintf(&#34;internal error: unknown goroutine (G %#x)&#34;, e.Goroutine)
<a id="L48"></a>}

<a id="L50"></a><span class="comment">// A NoCurrentGoroutine error occurs when no goroutine is currently</span>
<a id="L51"></a><span class="comment">// selected in a process (or when there are no goroutines in a</span>
<a id="L52"></a><span class="comment">// process).</span>
<a id="L53"></a>type NoCurrentGoroutine struct{}

<a id="L55"></a>func (e NoCurrentGoroutine) String() string { return &#34;no current goroutine&#34; }

<a id="L57"></a><span class="comment">// A Process represents a remote attached process.</span>
<a id="L58"></a>type Process struct {
    <a id="L59"></a>Arch;
    <a id="L60"></a>proc proc.Process;

    <a id="L62"></a><span class="comment">// The symbol table of this process</span>
    <a id="L63"></a>syms *gosym.Table;

    <a id="L65"></a><span class="comment">// A possibly-stopped OS thread, or nil</span>
    <a id="L66"></a>threadCache proc.Thread;

    <a id="L68"></a><span class="comment">// Types parsed from the remote process</span>
    <a id="L69"></a>types map[proc.Word]*remoteType;

    <a id="L71"></a><span class="comment">// Types and values from the remote runtime package</span>
    <a id="L72"></a>runtime runtimeValues;

    <a id="L74"></a><span class="comment">// Runtime field indexes</span>
    <a id="L75"></a>f   runtimeIndexes;

    <a id="L77"></a><span class="comment">// Globals from the sys package (or from no package)</span>
    <a id="L78"></a>sys struct {
        <a id="L79"></a>lessstack, goexit, newproc, deferproc, newprocreadylocked *gosym.Func;
        <a id="L80"></a>allg                                                      remotePtr;
        <a id="L81"></a>g0                                                        remoteStruct;
    <a id="L82"></a>};

    <a id="L84"></a><span class="comment">// Event queue</span>
    <a id="L85"></a>posted  []Event;
    <a id="L86"></a>pending []Event;
    <a id="L87"></a>event   Event;

    <a id="L89"></a><span class="comment">// Event hooks</span>
    <a id="L90"></a>breakpointHooks     map[proc.Word]*breakpointHook;
    <a id="L91"></a>goroutineCreateHook *goroutineCreateHook;
    <a id="L92"></a>goroutineExitHook   *goroutineExitHook;

    <a id="L94"></a><span class="comment">// Current goroutine, or nil if there are no goroutines</span>
    <a id="L95"></a>curGoroutine *Goroutine;

    <a id="L97"></a><span class="comment">// Goroutines by the address of their G structure</span>
    <a id="L98"></a>goroutines map[proc.Word]*Goroutine;
<a id="L99"></a>}

<a id="L101"></a><span class="comment">/*</span>
<a id="L102"></a><span class="comment"> * Process creation</span>
<a id="L103"></a><span class="comment"> */</span>

<a id="L105"></a><span class="comment">// NewProcess constructs a new remote process around a traced</span>
<a id="L106"></a><span class="comment">// process, an architecture, and a symbol table.</span>
<a id="L107"></a>func NewProcess(tproc proc.Process, arch Arch, syms *gosym.Table) (*Process, os.Error) {
    <a id="L108"></a>p := &amp;Process{
        <a id="L109"></a>Arch: arch,
        <a id="L110"></a>proc: tproc,
        <a id="L111"></a>syms: syms,
        <a id="L112"></a>types: make(map[proc.Word]*remoteType),
        <a id="L113"></a>breakpointHooks: make(map[proc.Word]*breakpointHook),
        <a id="L114"></a>goroutineCreateHook: new(goroutineCreateHook),
        <a id="L115"></a>goroutineExitHook: new(goroutineExitHook),
        <a id="L116"></a>goroutines: make(map[proc.Word]*Goroutine),
    <a id="L117"></a>};

    <a id="L119"></a><span class="comment">// Fill in remote runtime</span>
    <a id="L120"></a>p.bootstrap();

    <a id="L122"></a>switch {
    <a id="L123"></a>case p.sys.allg.addr().base == 0:
        <a id="L124"></a>return nil, FormatError(&#34;failed to find runtime symbol &#39;allg&#39;&#34;)
    <a id="L125"></a>case p.sys.g0.addr().base == 0:
        <a id="L126"></a>return nil, FormatError(&#34;failed to find runtime symbol &#39;g0&#39;&#34;)
    <a id="L127"></a>case p.sys.newprocreadylocked == nil:
        <a id="L128"></a>return nil, FormatError(&#34;failed to find runtime symbol &#39;newprocreadylocked&#39;&#34;)
    <a id="L129"></a>case p.sys.goexit == nil:
        <a id="L130"></a>return nil, FormatError(&#34;failed to find runtime symbol &#39;sys.goexit&#39;&#34;)
    <a id="L131"></a>}

    <a id="L133"></a><span class="comment">// Get current goroutines</span>
    <a id="L134"></a>p.goroutines[p.sys.g0.addr().base] = &amp;Goroutine{p.sys.g0, nil, false};
    <a id="L135"></a>err := try(func(a aborter) {
        <a id="L136"></a>g := p.sys.allg.aGet(a);
        <a id="L137"></a>for g != nil {
            <a id="L138"></a>gs := g.(remoteStruct);
            <a id="L139"></a>fmt.Printf(&#34;*** Found goroutine at %#x\n&#34;, gs.addr().base);
            <a id="L140"></a>p.goroutines[gs.addr().base] = &amp;Goroutine{gs, nil, false};
            <a id="L141"></a>g = gs.field(p.f.G.Alllink).(remotePtr).aGet(a);
        <a id="L142"></a>}
    <a id="L143"></a>});
    <a id="L144"></a>if err != nil {
        <a id="L145"></a>return nil, err
    <a id="L146"></a>}

    <a id="L148"></a><span class="comment">// Create internal breakpoints to catch new and exited goroutines</span>
    <a id="L149"></a>p.OnBreakpoint(proc.Word(p.sys.newprocreadylocked.Entry)).(*breakpointHook).addHandler(readylockedBP, true);
    <a id="L150"></a>p.OnBreakpoint(proc.Word(p.sys.goexit.Entry)).(*breakpointHook).addHandler(goexitBP, true);

    <a id="L152"></a><span class="comment">// Select current frames</span>
    <a id="L153"></a>for _, g := range p.goroutines {
        <a id="L154"></a>g.resetFrame()
    <a id="L155"></a>}

    <a id="L157"></a>p.selectSomeGoroutine();

    <a id="L159"></a>return p, nil;
<a id="L160"></a>}

<a id="L162"></a>func elfGoSyms(f *elf.File) (*gosym.Table, os.Error) {
    <a id="L163"></a>text := f.Section(&#34;.text&#34;);
    <a id="L164"></a>symtab := f.Section(&#34;.gosymtab&#34;);
    <a id="L165"></a>pclntab := f.Section(&#34;.gopclntab&#34;);
    <a id="L166"></a>if text == nil || symtab == nil || pclntab == nil {
        <a id="L167"></a>return nil, nil
    <a id="L168"></a>}

    <a id="L170"></a>symdat, err := symtab.Data();
    <a id="L171"></a>if err != nil {
        <a id="L172"></a>return nil, err
    <a id="L173"></a>}
    <a id="L174"></a>pclndat, err := pclntab.Data();
    <a id="L175"></a>if err != nil {
        <a id="L176"></a>return nil, err
    <a id="L177"></a>}

    <a id="L179"></a>pcln := gosym.NewLineTable(pclndat, text.Addr);
    <a id="L180"></a>tab, err := gosym.NewTable(symdat, pcln);
    <a id="L181"></a>if err != nil {
        <a id="L182"></a>return nil, err
    <a id="L183"></a>}

    <a id="L185"></a>return tab, nil;
<a id="L186"></a>}

<a id="L188"></a><span class="comment">// NewProcessElf constructs a new remote process around a traced</span>
<a id="L189"></a><span class="comment">// process and the process&#39; ELF object.</span>
<a id="L190"></a>func NewProcessElf(tproc proc.Process, f *elf.File) (*Process, os.Error) {
    <a id="L191"></a>syms, err := elfGoSyms(f);
    <a id="L192"></a>if err != nil {
        <a id="L193"></a>return nil, err
    <a id="L194"></a>}
    <a id="L195"></a>if syms == nil {
        <a id="L196"></a>return nil, FormatError(&#34;Failed to find symbol table&#34;)
    <a id="L197"></a>}
    <a id="L198"></a>var arch Arch;
    <a id="L199"></a>switch f.Machine {
    <a id="L200"></a>case elf.EM_X86_64:
        <a id="L201"></a>arch = Amd64
    <a id="L202"></a>default:
        <a id="L203"></a>return nil, UnknownArchitecture(f.Machine)
    <a id="L204"></a>}
    <a id="L205"></a>return NewProcess(tproc, arch, syms);
<a id="L206"></a>}

<a id="L208"></a><span class="comment">// bootstrap constructs the runtime structure of a remote process.</span>
<a id="L209"></a>func (p *Process) bootstrap() {
    <a id="L210"></a><span class="comment">// Manually construct runtime types</span>
    <a id="L211"></a>p.runtime.String = newManualType(eval.TypeOfNative(rt1String{}), p.Arch);
    <a id="L212"></a>p.runtime.Slice = newManualType(eval.TypeOfNative(rt1Slice{}), p.Arch);
    <a id="L213"></a>p.runtime.Eface = newManualType(eval.TypeOfNative(rt1Eface{}), p.Arch);

    <a id="L215"></a>p.runtime.Type = newManualType(eval.TypeOfNative(rt1Type{}), p.Arch);
    <a id="L216"></a>p.runtime.CommonType = newManualType(eval.TypeOfNative(rt1CommonType{}), p.Arch);
    <a id="L217"></a>p.runtime.UncommonType = newManualType(eval.TypeOfNative(rt1UncommonType{}), p.Arch);
    <a id="L218"></a>p.runtime.StructField = newManualType(eval.TypeOfNative(rt1StructField{}), p.Arch);
    <a id="L219"></a>p.runtime.StructType = newManualType(eval.TypeOfNative(rt1StructType{}), p.Arch);
    <a id="L220"></a>p.runtime.PtrType = newManualType(eval.TypeOfNative(rt1PtrType{}), p.Arch);
    <a id="L221"></a>p.runtime.ArrayType = newManualType(eval.TypeOfNative(rt1ArrayType{}), p.Arch);
    <a id="L222"></a>p.runtime.SliceType = newManualType(eval.TypeOfNative(rt1SliceType{}), p.Arch);

    <a id="L224"></a>p.runtime.Stktop = newManualType(eval.TypeOfNative(rt1Stktop{}), p.Arch);
    <a id="L225"></a>p.runtime.Gobuf = newManualType(eval.TypeOfNative(rt1Gobuf{}), p.Arch);
    <a id="L226"></a>p.runtime.G = newManualType(eval.TypeOfNative(rt1G{}), p.Arch);

    <a id="L228"></a><span class="comment">// Get addresses of type.*runtime.XType for discrimination.</span>
    <a id="L229"></a>rtv := reflect.Indirect(reflect.NewValue(&amp;p.runtime)).(*reflect.StructValue);
    <a id="L230"></a>rtvt := rtv.Type().(*reflect.StructType);
    <a id="L231"></a>for i := 0; i &lt; rtv.NumField(); i++ {
        <a id="L232"></a>n := rtvt.Field(i).Name;
        <a id="L233"></a>if n[0] != &#39;P&#39; || n[1] &lt; &#39;A&#39; || n[1] &gt; &#39;Z&#39; {
            <a id="L234"></a>continue
        <a id="L235"></a>}
        <a id="L236"></a>sym := p.syms.LookupSym(&#34;type.*runtime.&#34; + n[1:len(n)]);
        <a id="L237"></a>if sym == nil {
            <a id="L238"></a>continue
        <a id="L239"></a>}
        <a id="L240"></a>rtv.Field(i).(*reflect.Uint64Value).Set(sym.Value);
    <a id="L241"></a>}

    <a id="L243"></a><span class="comment">// Get runtime field indexes</span>
    <a id="L244"></a>fillRuntimeIndexes(&amp;p.runtime, &amp;p.f);

    <a id="L246"></a><span class="comment">// Fill G status</span>
    <a id="L247"></a>p.runtime.runtimeGStatus = rt1GStatus;

    <a id="L249"></a><span class="comment">// Get globals</span>
    <a id="L250"></a>p.sys.lessstack = p.syms.LookupFunc(&#34;sys.lessstack&#34;);
    <a id="L251"></a>p.sys.goexit = p.syms.LookupFunc(&#34;goexit&#34;);
    <a id="L252"></a>p.sys.newproc = p.syms.LookupFunc(&#34;sys.newproc&#34;);
    <a id="L253"></a>p.sys.deferproc = p.syms.LookupFunc(&#34;sys.deferproc&#34;);
    <a id="L254"></a>p.sys.newprocreadylocked = p.syms.LookupFunc(&#34;newprocreadylocked&#34;);
    <a id="L255"></a>if allg := p.syms.LookupSym(&#34;allg&#34;); allg != nil {
        <a id="L256"></a>p.sys.allg = remotePtr{remote{proc.Word(allg.Value), p}, p.runtime.G}
    <a id="L257"></a>}
    <a id="L258"></a>if g0 := p.syms.LookupSym(&#34;g0&#34;); g0 != nil {
        <a id="L259"></a>p.sys.g0 = p.runtime.G.mk(remote{proc.Word(g0.Value), p}).(remoteStruct)
    <a id="L260"></a>}
<a id="L261"></a>}

<a id="L263"></a>func (p *Process) selectSomeGoroutine() {
    <a id="L264"></a><span class="comment">// Once we have friendly goroutine ID&#39;s, there might be a more</span>
    <a id="L265"></a><span class="comment">// reasonable behavior for this.</span>
    <a id="L266"></a>p.curGoroutine = nil;
    <a id="L267"></a>for _, g := range p.goroutines {
        <a id="L268"></a>if !g.isG0() &amp;&amp; g.frame != nil {
            <a id="L269"></a>p.curGoroutine = g;
            <a id="L270"></a>return;
        <a id="L271"></a>}
    <a id="L272"></a>}
<a id="L273"></a>}

<a id="L275"></a><span class="comment">/*</span>
<a id="L276"></a><span class="comment"> * Process memory</span>
<a id="L277"></a><span class="comment"> */</span>

<a id="L279"></a>func (p *Process) someStoppedOSThread() proc.Thread {
    <a id="L280"></a>if p.threadCache != nil {
        <a id="L281"></a>if _, err := p.threadCache.Stopped(); err == nil {
            <a id="L282"></a>return p.threadCache
        <a id="L283"></a>}
    <a id="L284"></a>}

    <a id="L286"></a>for _, t := range p.proc.Threads() {
        <a id="L287"></a>if _, err := t.Stopped(); err == nil {
            <a id="L288"></a>p.threadCache = t;
            <a id="L289"></a>return t;
        <a id="L290"></a>}
    <a id="L291"></a>}
    <a id="L292"></a>return nil;
<a id="L293"></a>}

<a id="L295"></a>func (p *Process) Peek(addr proc.Word, out []byte) (int, os.Error) {
    <a id="L296"></a>thr := p.someStoppedOSThread();
    <a id="L297"></a>if thr == nil {
        <a id="L298"></a>return 0, ProcessNotStopped{}
    <a id="L299"></a>}
    <a id="L300"></a>return thr.Peek(addr, out);
<a id="L301"></a>}

<a id="L303"></a>func (p *Process) Poke(addr proc.Word, b []byte) (int, os.Error) {
    <a id="L304"></a>thr := p.someStoppedOSThread();
    <a id="L305"></a>if thr == nil {
        <a id="L306"></a>return 0, ProcessNotStopped{}
    <a id="L307"></a>}
    <a id="L308"></a>return thr.Poke(addr, b);
<a id="L309"></a>}

<a id="L311"></a>func (p *Process) peekUintptr(a aborter, addr proc.Word) proc.Word {
    <a id="L312"></a>return proc.Word(mkUintptr(remote{addr, p}).(remoteUint).aGet(a))
<a id="L313"></a>}

<a id="L315"></a><span class="comment">/*</span>
<a id="L316"></a><span class="comment"> * Events</span>
<a id="L317"></a><span class="comment"> */</span>

<a id="L319"></a><span class="comment">// OnBreakpoint returns the hook that is run when the program reaches</span>
<a id="L320"></a><span class="comment">// the given program counter.</span>
<a id="L321"></a>func (p *Process) OnBreakpoint(pc proc.Word) EventHook {
    <a id="L322"></a>if bp, ok := p.breakpointHooks[pc]; ok {
        <a id="L323"></a>return bp
    <a id="L324"></a>}
    <a id="L325"></a><span class="comment">// The breakpoint will register itself when a handler is added</span>
    <a id="L326"></a>return &amp;breakpointHook{commonHook{nil, 0}, p, pc};
<a id="L327"></a>}

<a id="L329"></a><span class="comment">// OnGoroutineCreate returns the hook that is run when a goroutine is created.</span>
<a id="L330"></a>func (p *Process) OnGoroutineCreate() EventHook {
    <a id="L331"></a>return p.goroutineCreateHook
<a id="L332"></a>}

<a id="L334"></a><span class="comment">// OnGoroutineExit returns the hook that is run when a goroutine exits.</span>
<a id="L335"></a>func (p *Process) OnGoroutineExit() EventHook { return p.goroutineExitHook }

<a id="L337"></a><span class="comment">// osThreadToGoroutine looks up the goroutine running on an OS thread.</span>
<a id="L338"></a>func (p *Process) osThreadToGoroutine(t proc.Thread) (*Goroutine, os.Error) {
    <a id="L339"></a>regs, err := t.Regs();
    <a id="L340"></a>if err != nil {
        <a id="L341"></a>return nil, err
    <a id="L342"></a>}
    <a id="L343"></a>g := p.G(regs);
    <a id="L344"></a>gt, ok := p.goroutines[g];
    <a id="L345"></a>if !ok {
        <a id="L346"></a>return nil, UnknownGoroutine{t, g}
    <a id="L347"></a>}
    <a id="L348"></a>return gt, nil;
<a id="L349"></a>}

<a id="L351"></a><span class="comment">// causesToEvents translates the stop causes of the underlying process</span>
<a id="L352"></a><span class="comment">// into an event queue.</span>
<a id="L353"></a>func (p *Process) causesToEvents() ([]Event, os.Error) {
    <a id="L354"></a><span class="comment">// Count causes we&#39;re interested in</span>
    <a id="L355"></a>nev := 0;
    <a id="L356"></a>for _, t := range p.proc.Threads() {
        <a id="L357"></a>if c, err := t.Stopped(); err == nil {
            <a id="L358"></a>switch c := c.(type) {
            <a id="L359"></a>case proc.Breakpoint:
                <a id="L360"></a>nev++
            <a id="L361"></a>case proc.Signal:
                <a id="L362"></a><span class="comment">// TODO(austin)</span>
                <a id="L363"></a><span class="comment">//nev++;</span>
            <a id="L364"></a>}
        <a id="L365"></a>}
    <a id="L366"></a>}

    <a id="L368"></a><span class="comment">// Translate causes to events</span>
    <a id="L369"></a>events := make([]Event, nev);
    <a id="L370"></a>i := 0;
    <a id="L371"></a>for _, t := range p.proc.Threads() {
        <a id="L372"></a>if c, err := t.Stopped(); err == nil {
            <a id="L373"></a>switch c := c.(type) {
            <a id="L374"></a>case proc.Breakpoint:
                <a id="L375"></a>gt, err := p.osThreadToGoroutine(t);
                <a id="L376"></a>if err != nil {
                    <a id="L377"></a>return nil, err
                <a id="L378"></a>}
                <a id="L379"></a>events[i] = &amp;Breakpoint{commonEvent{p, gt}, t, proc.Word(c)};
                <a id="L380"></a>i++;
            <a id="L381"></a>case proc.Signal:
                <a id="L382"></a><span class="comment">// TODO(austin)</span>
            <a id="L383"></a>}
        <a id="L384"></a>}
    <a id="L385"></a>}

    <a id="L387"></a>return events, nil;
<a id="L388"></a>}

<a id="L390"></a><span class="comment">// postEvent appends an event to the posted queue.  These events will</span>
<a id="L391"></a><span class="comment">// be processed before any currently pending events.</span>
<a id="L392"></a>func (p *Process) postEvent(ev Event) {
    <a id="L393"></a>n := len(p.posted);
    <a id="L394"></a>m := n * 2;
    <a id="L395"></a>if m == 0 {
        <a id="L396"></a>m = 4
    <a id="L397"></a>}
    <a id="L398"></a>posted := make([]Event, n+1, m);
    <a id="L399"></a>for i, p := range p.posted {
        <a id="L400"></a>posted[i] = p
    <a id="L401"></a>}
    <a id="L402"></a>posted[n] = ev;
    <a id="L403"></a>p.posted = posted;
<a id="L404"></a>}

<a id="L406"></a><span class="comment">// processEvents processes events in the event queue until no events</span>
<a id="L407"></a><span class="comment">// remain, a handler returns EAStop, or a handler returns an error.</span>
<a id="L408"></a><span class="comment">// It returns either EAStop or EAContinue and possibly an error.</span>
<a id="L409"></a>func (p *Process) processEvents() (EventAction, os.Error) {
    <a id="L410"></a>var ev Event;
    <a id="L411"></a>for len(p.posted) &gt; 0 {
        <a id="L412"></a>ev, p.posted = p.posted[0], p.posted[1:len(p.posted)];
        <a id="L413"></a>action, err := p.processEvent(ev);
        <a id="L414"></a>if action == EAStop {
            <a id="L415"></a>return action, err
        <a id="L416"></a>}
    <a id="L417"></a>}

    <a id="L419"></a>for len(p.pending) &gt; 0 {
        <a id="L420"></a>ev, p.pending = p.pending[0], p.pending[1:len(p.pending)];
        <a id="L421"></a>action, err := p.processEvent(ev);
        <a id="L422"></a>if action == EAStop {
            <a id="L423"></a>return action, err
        <a id="L424"></a>}
    <a id="L425"></a>}

    <a id="L427"></a>return EAContinue, nil;
<a id="L428"></a>}

<a id="L430"></a><span class="comment">// processEvent processes a single event, without manipulating the</span>
<a id="L431"></a><span class="comment">// event queues.  It returns either EAStop or EAContinue and possibly</span>
<a id="L432"></a><span class="comment">// an error.</span>
<a id="L433"></a>func (p *Process) processEvent(ev Event) (EventAction, os.Error) {
    <a id="L434"></a>p.event = ev;

    <a id="L436"></a>var action EventAction;
    <a id="L437"></a>var err os.Error;
    <a id="L438"></a>switch ev := p.event.(type) {
    <a id="L439"></a>case *Breakpoint:
        <a id="L440"></a>hook, ok := p.breakpointHooks[ev.pc];
        <a id="L441"></a>if !ok {
            <a id="L442"></a>break
        <a id="L443"></a>}
        <a id="L444"></a>p.curGoroutine = ev.Goroutine();
        <a id="L445"></a>action, err = hook.handle(ev);

    <a id="L447"></a>case *GoroutineCreate:
        <a id="L448"></a>p.curGoroutine = ev.Goroutine();
        <a id="L449"></a>action, err = p.goroutineCreateHook.handle(ev);

    <a id="L451"></a>case *GoroutineExit:
        <a id="L452"></a>action, err = p.goroutineExitHook.handle(ev)

    <a id="L454"></a>default:
        <a id="L455"></a>log.Crashf(&#34;Unknown event type %T in queue&#34;, p.event)
    <a id="L456"></a>}

    <a id="L458"></a>if err != nil {
        <a id="L459"></a>return EAStop, err
    <a id="L460"></a>} else if action == EAStop {
        <a id="L461"></a>return EAStop, nil
    <a id="L462"></a>}
    <a id="L463"></a>return EAContinue, nil;
<a id="L464"></a>}

<a id="L466"></a><span class="comment">// Event returns the last event that caused the process to stop.  This</span>
<a id="L467"></a><span class="comment">// may return nil if the process has never been stopped by an event.</span>
<a id="L468"></a><span class="comment">//</span>
<a id="L469"></a><span class="comment">// TODO(austin) Return nil if the user calls p.Stop()?</span>
<a id="L470"></a>func (p *Process) Event() Event { return p.event }

<a id="L472"></a><span class="comment">/*</span>
<a id="L473"></a><span class="comment"> * Process control</span>
<a id="L474"></a><span class="comment"> */</span>

<a id="L476"></a><span class="comment">// TODO(austin) Cont, WaitStop, and Stop.  Need to figure out how</span>
<a id="L477"></a><span class="comment">// event handling works with these.  Originally I did it only in</span>
<a id="L478"></a><span class="comment">// WaitStop, but if you Cont and there are pending events, then you</span>
<a id="L479"></a><span class="comment">// have to not actually continue and wait until a WaitStop to process</span>
<a id="L480"></a><span class="comment">// them, even if the event handlers will tell you to continue.  We</span>
<a id="L481"></a><span class="comment">// could handle them in both Cont and WaitStop to avoid this problem,</span>
<a id="L482"></a><span class="comment">// but it&#39;s still weird if an event happens after the Cont and before</span>
<a id="L483"></a><span class="comment">// the WaitStop that the handlers say to continue from.  Or we could</span>
<a id="L484"></a><span class="comment">// handle them on a separate thread.  Then obviously you get weird</span>
<a id="L485"></a><span class="comment">// asynchronous things, like prints while the user it typing a command,</span>
<a id="L486"></a><span class="comment">// but that&#39;s not necessarily a bad thing.</span>

<a id="L488"></a><span class="comment">// ContWait resumes process execution and waits for an event to occur</span>
<a id="L489"></a><span class="comment">// that stops the process.</span>
<a id="L490"></a>func (p *Process) ContWait() os.Error {
    <a id="L491"></a>for {
        <a id="L492"></a>a, err := p.processEvents();
        <a id="L493"></a>if err != nil {
            <a id="L494"></a>return err
        <a id="L495"></a>} else if a == EAStop {
            <a id="L496"></a>break
        <a id="L497"></a>}
        <a id="L498"></a>err = p.proc.Continue();
        <a id="L499"></a>if err != nil {
            <a id="L500"></a>return err
        <a id="L501"></a>}
        <a id="L502"></a>err = p.proc.WaitStop();
        <a id="L503"></a>if err != nil {
            <a id="L504"></a>return err
        <a id="L505"></a>}
        <a id="L506"></a>for _, g := range p.goroutines {
            <a id="L507"></a>g.resetFrame()
        <a id="L508"></a>}
        <a id="L509"></a>p.pending, err = p.causesToEvents();
        <a id="L510"></a>if err != nil {
            <a id="L511"></a>return err
        <a id="L512"></a>}
    <a id="L513"></a>}
    <a id="L514"></a>return nil;
<a id="L515"></a>}

<a id="L517"></a><span class="comment">// Out selects the caller frame of the current frame.</span>
<a id="L518"></a>func (p *Process) Out() os.Error {
    <a id="L519"></a>if p.curGoroutine == nil {
        <a id="L520"></a>return NoCurrentGoroutine{}
    <a id="L521"></a>}
    <a id="L522"></a>return p.curGoroutine.Out();
<a id="L523"></a>}

<a id="L525"></a><span class="comment">// In selects the frame called by the current frame.</span>
<a id="L526"></a>func (p *Process) In() os.Error {
    <a id="L527"></a>if p.curGoroutine == nil {
        <a id="L528"></a>return NoCurrentGoroutine{}
    <a id="L529"></a>}
    <a id="L530"></a>return p.curGoroutine.In();
<a id="L531"></a>}
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
