<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/ogle/cmd.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/exp/ogle/cmd.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Ogle is the beginning of a debugger for Go.</span>
<a id="L6"></a>package ogle

<a id="L8"></a>import (
    <a id="L9"></a>&#34;bufio&#34;;
    <a id="L10"></a>&#34;debug/elf&#34;;
    <a id="L11"></a>&#34;debug/proc&#34;;
    <a id="L12"></a>&#34;exp/eval&#34;;
    <a id="L13"></a>&#34;fmt&#34;;
    <a id="L14"></a>&#34;go/scanner&#34;;
    <a id="L15"></a>&#34;go/token&#34;;
    <a id="L16"></a>&#34;os&#34;;
    <a id="L17"></a>&#34;strconv&#34;;
    <a id="L18"></a>&#34;strings&#34;;
<a id="L19"></a>)

<a id="L21"></a>var world *eval.World
<a id="L22"></a>var curProc *Process

<a id="L24"></a>func Main() {
    <a id="L25"></a>world = eval.NewWorld();
    <a id="L26"></a>defineFuncs();
    <a id="L27"></a>r := bufio.NewReader(os.Stdin);
    <a id="L28"></a>for {
        <a id="L29"></a>print(&#34;; &#34;);
        <a id="L30"></a>line, err := r.ReadSlice(&#39;\n&#39;);
        <a id="L31"></a>if err != nil {
            <a id="L32"></a>break
        <a id="L33"></a>}

        <a id="L35"></a><span class="comment">// Try line as a command</span>
        <a id="L36"></a>cmd, rest := getCmd(line);
        <a id="L37"></a>if cmd != nil {
            <a id="L38"></a>err := cmd.handler(rest);
            <a id="L39"></a>if err != nil {
                <a id="L40"></a>scanner.PrintError(os.Stderr, err)
            <a id="L41"></a>}
            <a id="L42"></a>continue;
        <a id="L43"></a>}

        <a id="L45"></a><span class="comment">// Try line as code</span>
        <a id="L46"></a>code, err := world.Compile(string(line));
        <a id="L47"></a>if err != nil {
            <a id="L48"></a>scanner.PrintError(os.Stderr, err);
            <a id="L49"></a>continue;
        <a id="L50"></a>}
        <a id="L51"></a>v, err := code.Run();
        <a id="L52"></a>if err != nil {
            <a id="L53"></a>fmt.Fprintf(os.Stderr, err.String());
            <a id="L54"></a>continue;
        <a id="L55"></a>}
        <a id="L56"></a>if v != nil {
            <a id="L57"></a>println(v.String())
        <a id="L58"></a>}
    <a id="L59"></a>}
<a id="L60"></a>}

<a id="L62"></a><span class="comment">// newScanner creates a new scanner that scans that given input bytes.</span>
<a id="L63"></a>func newScanner(input []byte) (*scanner.Scanner, *scanner.ErrorVector) {
    <a id="L64"></a>sc := new(scanner.Scanner);
    <a id="L65"></a>ev := new(scanner.ErrorVector);
    <a id="L66"></a>ev.Init();
    <a id="L67"></a>sc.Init(&#34;input&#34;, input, ev, 0);

    <a id="L69"></a>return sc, ev;
<a id="L70"></a>}

<a id="L72"></a><span class="comment">/*</span>
<a id="L73"></a><span class="comment"> * Commands</span>
<a id="L74"></a><span class="comment"> */</span>

<a id="L76"></a><span class="comment">// A UsageError occurs when a command is called with illegal arguments.</span>
<a id="L77"></a>type UsageError string

<a id="L79"></a>func (e UsageError) String() string { return string(e) }

<a id="L81"></a><span class="comment">// A cmd represents a single command with a handler.</span>
<a id="L82"></a>type cmd struct {
    <a id="L83"></a>cmd     string;
    <a id="L84"></a>handler func([]byte) os.Error;
<a id="L85"></a>}

<a id="L87"></a>var cmds = []cmd{
    <a id="L88"></a>cmd{&#34;load&#34;, cmdLoad},
    <a id="L89"></a>cmd{&#34;bt&#34;, cmdBt},
<a id="L90"></a>}

<a id="L92"></a><span class="comment">// getCmd attempts to parse an input line as a registered command.  If</span>
<a id="L93"></a><span class="comment">// successful, it returns the command and the bytes remaining after</span>
<a id="L94"></a><span class="comment">// the command, which should be passed to the command.</span>
<a id="L95"></a>func getCmd(line []byte) (*cmd, []byte) {
    <a id="L96"></a>sc, _ := newScanner(line);
    <a id="L97"></a>pos, tok, lit := sc.Scan();
    <a id="L98"></a>if sc.ErrorCount != 0 || tok != token.IDENT {
        <a id="L99"></a>return nil, nil
    <a id="L100"></a>}

    <a id="L102"></a>slit := string(lit);
    <a id="L103"></a>for i := range cmds {
        <a id="L104"></a>if cmds[i].cmd == slit {
            <a id="L105"></a>return &amp;cmds[i], line[pos.Offset+len(lit) : len(line)]
        <a id="L106"></a>}
    <a id="L107"></a>}
    <a id="L108"></a>return nil, nil;
<a id="L109"></a>}

<a id="L111"></a><span class="comment">// cmdLoad starts or attaches to a process.  Its form is similar to</span>
<a id="L112"></a><span class="comment">// import:</span>
<a id="L113"></a><span class="comment">//</span>
<a id="L114"></a><span class="comment">//  load [sym] &#34;path&#34; [;]</span>
<a id="L115"></a><span class="comment">//</span>
<a id="L116"></a><span class="comment">// sym specifies the name to give to the process.  If not given, the</span>
<a id="L117"></a><span class="comment">// name is derived from the path of the process.  If &#34;.&#34;, then the</span>
<a id="L118"></a><span class="comment">// packages from the remote process are defined into the current</span>
<a id="L119"></a><span class="comment">// namespace.  If given, this symbol is defined as a package</span>
<a id="L120"></a><span class="comment">// containing the process&#39; packages.</span>
<a id="L121"></a><span class="comment">//</span>
<a id="L122"></a><span class="comment">// path gives the path of the process to start or attach to.  If it is</span>
<a id="L123"></a><span class="comment">// &#34;pid:&lt;num&gt;&#34;, then attach to the given PID.  Otherwise, treat it as</span>
<a id="L124"></a><span class="comment">// a file path and space-separated arguments and start a new process.</span>
<a id="L125"></a><span class="comment">//</span>
<a id="L126"></a><span class="comment">// load always sets the current process to the loaded process.</span>
<a id="L127"></a>func cmdLoad(args []byte) os.Error {
    <a id="L128"></a>ident, path, err := parseLoad(args);
    <a id="L129"></a>if err != nil {
        <a id="L130"></a>return err
    <a id="L131"></a>}
    <a id="L132"></a>if curProc != nil {
        <a id="L133"></a>return UsageError(&#34;multiple processes not implemented&#34;)
    <a id="L134"></a>}
    <a id="L135"></a>if ident != &#34;.&#34; {
        <a id="L136"></a>return UsageError(&#34;process identifiers not implemented&#34;)
    <a id="L137"></a>}

    <a id="L139"></a><span class="comment">// Parse argument and start or attach to process</span>
    <a id="L140"></a>var fname string;
    <a id="L141"></a>var tproc proc.Process;
    <a id="L142"></a>if len(path) &gt;= 4 &amp;&amp; path[0:4] == &#34;pid:&#34; {
        <a id="L143"></a>pid, err := strconv.Atoi(path[4:len(path)]);
        <a id="L144"></a>if err != nil {
            <a id="L145"></a>return err
        <a id="L146"></a>}
        <a id="L147"></a>fname, err = os.Readlink(fmt.Sprintf(&#34;/proc/%d/exe&#34;, pid));
        <a id="L148"></a>if err != nil {
            <a id="L149"></a>return err
        <a id="L150"></a>}
        <a id="L151"></a>tproc, err = proc.Attach(pid);
        <a id="L152"></a>if err != nil {
            <a id="L153"></a>return err
        <a id="L154"></a>}
        <a id="L155"></a>println(&#34;Attached to&#34;, pid);
    <a id="L156"></a>} else {
        <a id="L157"></a>parts := strings.Split(path, &#34; &#34;, 0);
        <a id="L158"></a>if len(parts) == 0 {
            <a id="L159"></a>fname = &#34;&#34;
        <a id="L160"></a>} else {
            <a id="L161"></a>fname = parts[0]
        <a id="L162"></a>}
        <a id="L163"></a>tproc, err = proc.ForkExec(fname, parts, os.Environ(), &#34;&#34;, []*os.File{os.Stdin, os.Stdout, os.Stderr});
        <a id="L164"></a>if err != nil {
            <a id="L165"></a>return err
        <a id="L166"></a>}
        <a id="L167"></a>println(&#34;Started&#34;, path);
        <a id="L168"></a><span class="comment">// TODO(austin) If we fail after this point, kill tproc</span>
        <a id="L169"></a><span class="comment">// before detaching.</span>
    <a id="L170"></a>}

    <a id="L172"></a><span class="comment">// Get symbols</span>
    <a id="L173"></a>f, err := os.Open(fname, os.O_RDONLY, 0);
    <a id="L174"></a>if err != nil {
        <a id="L175"></a>tproc.Detach();
        <a id="L176"></a>return err;
    <a id="L177"></a>}
    <a id="L178"></a>defer f.Close();
    <a id="L179"></a>elf, err := elf.NewFile(f);
    <a id="L180"></a>if err != nil {
        <a id="L181"></a>tproc.Detach();
        <a id="L182"></a>return err;
    <a id="L183"></a>}
    <a id="L184"></a>curProc, err = NewProcessElf(tproc, elf);
    <a id="L185"></a>if err != nil {
        <a id="L186"></a>tproc.Detach();
        <a id="L187"></a>return err;
    <a id="L188"></a>}

    <a id="L190"></a><span class="comment">// Prepare new process</span>
    <a id="L191"></a>curProc.OnGoroutineCreate().AddHandler(EventPrint);
    <a id="L192"></a>curProc.OnGoroutineExit().AddHandler(EventPrint);

    <a id="L194"></a>err = curProc.populateWorld(world);
    <a id="L195"></a>if err != nil {
        <a id="L196"></a>tproc.Detach();
        <a id="L197"></a>return err;
    <a id="L198"></a>}

    <a id="L200"></a>return nil;
<a id="L201"></a>}

<a id="L203"></a>func parseLoad(args []byte) (ident string, path string, err os.Error) {
    <a id="L204"></a>err = UsageError(&#34;Usage: load [sym] \&#34;path\&#34;&#34;);
    <a id="L205"></a>sc, ev := newScanner(args);

    <a id="L207"></a>var toks [4]token.Token;
    <a id="L208"></a>var lits [4][]byte;
    <a id="L209"></a>for i := range toks {
        <a id="L210"></a>_, toks[i], lits[i] = sc.Scan()
    <a id="L211"></a>}
    <a id="L212"></a>if sc.ErrorCount != 0 {
        <a id="L213"></a>err = ev.GetError(scanner.NoMultiples);
        <a id="L214"></a>return;
    <a id="L215"></a>}

    <a id="L217"></a>i := 0;
    <a id="L218"></a>switch toks[i] {
    <a id="L219"></a>case token.PERIOD, token.IDENT:
        <a id="L220"></a>ident = string(lits[i]);
        <a id="L221"></a>i++;
    <a id="L222"></a>}

    <a id="L224"></a>if toks[i] != token.STRING {
        <a id="L225"></a>return
    <a id="L226"></a>}
    <a id="L227"></a>path, uerr := strconv.Unquote(string(lits[i]));
    <a id="L228"></a>if uerr != nil {
        <a id="L229"></a>err = uerr;
        <a id="L230"></a>return;
    <a id="L231"></a>}
    <a id="L232"></a>i++;

    <a id="L234"></a>if toks[i] == token.SEMICOLON {
        <a id="L235"></a>i++
    <a id="L236"></a>}
    <a id="L237"></a>if toks[i] != token.EOF {
        <a id="L238"></a>return
    <a id="L239"></a>}

    <a id="L241"></a>return ident, path, nil;
<a id="L242"></a>}

<a id="L244"></a><span class="comment">// cmdBt prints a backtrace for the current goroutine.  It takes no</span>
<a id="L245"></a><span class="comment">// arguments.</span>
<a id="L246"></a>func cmdBt(args []byte) os.Error {
    <a id="L247"></a>err := parseNoArgs(args, &#34;Usage: bt&#34;);
    <a id="L248"></a>if err != nil {
        <a id="L249"></a>return err
    <a id="L250"></a>}

    <a id="L252"></a>if curProc == nil || curProc.curGoroutine == nil {
        <a id="L253"></a>return NoCurrentGoroutine{}
    <a id="L254"></a>}

    <a id="L256"></a>f := curProc.curGoroutine.frame;
    <a id="L257"></a>if f == nil {
        <a id="L258"></a>fmt.Println(&#34;No frames on stack&#34;);
        <a id="L259"></a>return nil;
    <a id="L260"></a>}

    <a id="L262"></a>for f.Inner() != nil {
        <a id="L263"></a>f = f.Inner()
    <a id="L264"></a>}

    <a id="L266"></a>for i := 0; i &lt; 100; i++ {
        <a id="L267"></a>if f == curProc.curGoroutine.frame {
            <a id="L268"></a>fmt.Printf(&#34;=&gt; &#34;)
        <a id="L269"></a>} else {
            <a id="L270"></a>fmt.Printf(&#34;   &#34;)
        <a id="L271"></a>}
        <a id="L272"></a>fmt.Printf(&#34;%8x %v\n&#34;, f.pc, f);
        <a id="L273"></a>f, err = f.Outer();
        <a id="L274"></a>if err != nil {
            <a id="L275"></a>return err
        <a id="L276"></a>}
        <a id="L277"></a>if f == nil {
            <a id="L278"></a>return nil
        <a id="L279"></a>}
    <a id="L280"></a>}

    <a id="L282"></a>fmt.Println(&#34;...&#34;);
    <a id="L283"></a>return nil;
<a id="L284"></a>}

<a id="L286"></a>func parseNoArgs(args []byte, usage string) os.Error {
    <a id="L287"></a>sc, ev := newScanner(args);
    <a id="L288"></a>_, tok, _ := sc.Scan();
    <a id="L289"></a>if sc.ErrorCount != 0 {
        <a id="L290"></a>return ev.GetError(scanner.NoMultiples)
    <a id="L291"></a>}
    <a id="L292"></a>if tok != token.EOF {
        <a id="L293"></a>return UsageError(usage)
    <a id="L294"></a>}
    <a id="L295"></a>return nil;
<a id="L296"></a>}

<a id="L298"></a><span class="comment">/*</span>
<a id="L299"></a><span class="comment"> * Functions</span>
<a id="L300"></a><span class="comment"> */</span>

<a id="L302"></a><span class="comment">// defineFuncs populates world with the built-in functions.</span>
<a id="L303"></a>func defineFuncs() {
    <a id="L304"></a>t, v := eval.FuncFromNativeTyped(fnOut, fnOutSig);
    <a id="L305"></a>world.DefineConst(&#34;Out&#34;, t, v);
    <a id="L306"></a>t, v = eval.FuncFromNativeTyped(fnContWait, fnContWaitSig);
    <a id="L307"></a>world.DefineConst(&#34;ContWait&#34;, t, v);
    <a id="L308"></a>t, v = eval.FuncFromNativeTyped(fnBpSet, fnBpSetSig);
    <a id="L309"></a>world.DefineConst(&#34;BpSet&#34;, t, v);
<a id="L310"></a>}

<a id="L312"></a><span class="comment">// printCurFrame prints the current stack frame, as it would appear in</span>
<a id="L313"></a><span class="comment">// a backtrace.</span>
<a id="L314"></a>func printCurFrame() {
    <a id="L315"></a>if curProc == nil || curProc.curGoroutine == nil {
        <a id="L316"></a>return
    <a id="L317"></a>}
    <a id="L318"></a>f := curProc.curGoroutine.frame;
    <a id="L319"></a>if f == nil {
        <a id="L320"></a>return
    <a id="L321"></a>}
    <a id="L322"></a>fmt.Printf(&#34;=&gt; %8x %v\n&#34;, f.pc, f);
<a id="L323"></a>}

<a id="L325"></a><span class="comment">// fnOut moves the current frame to the caller of the current frame.</span>
<a id="L326"></a>func fnOutSig() {}
<a id="L327"></a>func fnOut(t *eval.Thread, args []eval.Value, res []eval.Value) {
    <a id="L328"></a>if curProc == nil {
        <a id="L329"></a>t.Abort(NoCurrentGoroutine{})
    <a id="L330"></a>}
    <a id="L331"></a>err := curProc.Out();
    <a id="L332"></a>if err != nil {
        <a id="L333"></a>t.Abort(err)
    <a id="L334"></a>}
    <a id="L335"></a><span class="comment">// TODO(austin) Only in the command form</span>
    <a id="L336"></a>printCurFrame();
<a id="L337"></a>}

<a id="L339"></a><span class="comment">// fnContWait continues the current process and waits for a stopping event.</span>
<a id="L340"></a>func fnContWaitSig() {}
<a id="L341"></a>func fnContWait(t *eval.Thread, args []eval.Value, res []eval.Value) {
    <a id="L342"></a>if curProc == nil {
        <a id="L343"></a>t.Abort(NoCurrentGoroutine{})
    <a id="L344"></a>}
    <a id="L345"></a>err := curProc.ContWait();
    <a id="L346"></a>if err != nil {
        <a id="L347"></a>t.Abort(err)
    <a id="L348"></a>}
    <a id="L349"></a><span class="comment">// TODO(austin) Only in the command form</span>
    <a id="L350"></a>ev := curProc.Event();
    <a id="L351"></a>if ev != nil {
        <a id="L352"></a>fmt.Printf(&#34;%v\n&#34;, ev)
    <a id="L353"></a>}
    <a id="L354"></a>printCurFrame();
<a id="L355"></a>}

<a id="L357"></a><span class="comment">// fnBpSet sets a breakpoint at the entry to the named function.</span>
<a id="L358"></a>func fnBpSetSig(string) {}
<a id="L359"></a>func fnBpSet(t *eval.Thread, args []eval.Value, res []eval.Value) {
    <a id="L360"></a><span class="comment">// TODO(austin) This probably shouldn&#39;t take a symbol name.</span>
    <a id="L361"></a><span class="comment">// Perhaps it should take an interface that provides PC&#39;s.</span>
    <a id="L362"></a><span class="comment">// Functions and instructions can implement that interface and</span>
    <a id="L363"></a><span class="comment">// we can have something to translate file:line pairs.</span>
    <a id="L364"></a>if curProc == nil {
        <a id="L365"></a>t.Abort(NoCurrentGoroutine{})
    <a id="L366"></a>}
    <a id="L367"></a>name := args[0].(eval.StringValue).Get(t);
    <a id="L368"></a>fn := curProc.syms.LookupFunc(name);
    <a id="L369"></a>if fn == nil {
        <a id="L370"></a>t.Abort(UsageError(&#34;no such function &#34; + name))
    <a id="L371"></a>}
    <a id="L372"></a>curProc.OnBreakpoint(proc.Word(fn.Entry)).AddHandler(EventStop);
<a id="L373"></a>}
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
