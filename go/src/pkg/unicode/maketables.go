<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/unicode/maketables.go</title>

  <link rel="stylesheet" type="text/css" href="../../../doc/style.css">
  <script type="text/javascript" src="../../../doc/godocs.js"></script>

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
        <a href="../../../index.html"><img src="../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../index.html">Source files</a></li>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/unicode/maketables.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Unicode table generator.</span>
<a id="L6"></a><span class="comment">// Data read from the web.</span>

<a id="L8"></a>package main

<a id="L10"></a>import (
    <a id="L11"></a>&#34;bufio&#34;;
    <a id="L12"></a>&#34;flag&#34;;
    <a id="L13"></a>&#34;fmt&#34;;
    <a id="L14"></a>&#34;http&#34;;
    <a id="L15"></a>&#34;log&#34;;
    <a id="L16"></a>&#34;os&#34;;
    <a id="L17"></a>&#34;sort&#34;;
    <a id="L18"></a>&#34;strconv&#34;;
    <a id="L19"></a>&#34;strings&#34;;
    <a id="L20"></a>&#34;regexp&#34;;
    <a id="L21"></a>&#34;unicode&#34;;
<a id="L22"></a>)

<a id="L24"></a>func main() {
    <a id="L25"></a>flag.Parse();
    <a id="L26"></a>loadChars(); <span class="comment">// always needed</span>
    <a id="L27"></a>printCategories();
    <a id="L28"></a>printScriptOrProperty(false);
    <a id="L29"></a>printScriptOrProperty(true);
    <a id="L30"></a>printCases();
<a id="L31"></a>}

<a id="L33"></a>var dataURL = flag.String(&#34;data&#34;, &#34;&#34;, &#34;full URL for UnicodeData.txt; defaults to --url/UnicodeData.txt&#34;)
<a id="L34"></a>var url = flag.String(&#34;url&#34;,
    <a id="L35"></a>&#34;http://www.unicode.org/Public/5.1.0/ucd/&#34;,
    <a id="L36"></a>&#34;URL of Unicode database directory&#34;)
<a id="L37"></a>var tablelist = flag.String(&#34;tables&#34;,
    <a id="L38"></a>&#34;all&#34;,
    <a id="L39"></a>&#34;comma-separated list of which tables to generate; can be letter&#34;)
<a id="L40"></a>var scriptlist = flag.String(&#34;scripts&#34;,
    <a id="L41"></a>&#34;all&#34;,
    <a id="L42"></a>&#34;comma-separated list of which script tables to generate&#34;)
<a id="L43"></a>var proplist = flag.String(&#34;props&#34;,
    <a id="L44"></a>&#34;all&#34;,
    <a id="L45"></a>&#34;comma-separated list of which property tables to generate&#34;)
<a id="L46"></a>var cases = flag.Bool(&#34;cases&#34;,
    <a id="L47"></a>true,
    <a id="L48"></a>&#34;generate case tables&#34;)
<a id="L49"></a>var test = flag.Bool(&#34;test&#34;,
    <a id="L50"></a>false,
    <a id="L51"></a>&#34;test existing tables; can be used to compare web data with package data&#34;)

<a id="L53"></a>var scriptRe = regexp.MustCompile(`([0-9A-F]+)(\.\.[0-9A-F]+)? *; ([A-Za-z_]+)`)
<a id="L54"></a>var die = log.New(os.Stderr, nil, &#34;&#34;, log.Lexit|log.Lshortfile)

<a id="L56"></a>var category = map[string]bool{&#34;letter&#34;: true} <span class="comment">// Nd Lu etc. letter is a special case</span>

<a id="L58"></a><span class="comment">// UnicodeData.txt has form:</span>
<a id="L59"></a><span class="comment">//	0037;DIGIT SEVEN;Nd;0;EN;;7;7;7;N;;;;;</span>
<a id="L60"></a><span class="comment">//	007A;LATIN SMALL LETTER Z;Ll;0;L;;;;;N;;;005A;;005A</span>
<a id="L61"></a><span class="comment">// See http://www.unicode.org/Public/5.1.0/ucd/UCD.html for full explanation</span>
<a id="L62"></a><span class="comment">// The fields:</span>
<a id="L63"></a>const (
    <a id="L64"></a>FCodePoint = iota;
    <a id="L65"></a>FName;
    <a id="L66"></a>FGeneralCategory;
    <a id="L67"></a>FCanonicalCombiningClass;
    <a id="L68"></a>FBidiClass;
    <a id="L69"></a>FDecompositionType;
    <a id="L70"></a>FDecompositionMapping;
    <a id="L71"></a>FNumericType;
    <a id="L72"></a>FNumericValue;
    <a id="L73"></a>FBidiMirrored;
    <a id="L74"></a>FUnicode1Name;
    <a id="L75"></a>FISOComment;
    <a id="L76"></a>FSimpleUppercaseMapping;
    <a id="L77"></a>FSimpleLowercaseMapping;
    <a id="L78"></a>FSimpleTitlecaseMapping;
    <a id="L79"></a>NumField;

    <a id="L81"></a>MaxChar = 0x10FFFF; <span class="comment">// anything above this shouldn&#39;t exist</span>
<a id="L82"></a>)

<a id="L84"></a>var fieldName = []string{
    <a id="L85"></a>&#34;CodePoint&#34;,
    <a id="L86"></a>&#34;Name&#34;,
    <a id="L87"></a>&#34;GeneralCategory&#34;,
    <a id="L88"></a>&#34;CanonicalCombiningClass&#34;,
    <a id="L89"></a>&#34;BidiClass&#34;,
    <a id="L90"></a>&#34;DecompositionType&#34;,
    <a id="L91"></a>&#34;DecompositionMapping&#34;,
    <a id="L92"></a>&#34;NumericType&#34;,
    <a id="L93"></a>&#34;NumericValue&#34;,
    <a id="L94"></a>&#34;BidiMirrored&#34;,
    <a id="L95"></a>&#34;Unicode1Name&#34;,
    <a id="L96"></a>&#34;ISOComment&#34;,
    <a id="L97"></a>&#34;SimpleUppercaseMapping&#34;,
    <a id="L98"></a>&#34;SimpleLowercaseMapping&#34;,
    <a id="L99"></a>&#34;SimpleTitlecaseMapping&#34;,
<a id="L100"></a>}

<a id="L102"></a><span class="comment">// This contains only the properties we&#39;re interested in.</span>
<a id="L103"></a>type Char struct {
    <a id="L104"></a>field     []string; <span class="comment">// debugging only; could be deleted if we take out char.dump()</span>
    <a id="L105"></a>codePoint uint32;   <span class="comment">// if zero, this index is not a valid code point.</span>
    <a id="L106"></a>category  string;
    <a id="L107"></a>upperCase int;
    <a id="L108"></a>lowerCase int;
    <a id="L109"></a>titleCase int;
<a id="L110"></a>}

<a id="L112"></a><span class="comment">// Scripts.txt has form:</span>
<a id="L113"></a><span class="comment">//	A673          ; Cyrillic # Po       SLAVONIC ASTERISK</span>
<a id="L114"></a><span class="comment">//	A67C..A67D    ; Cyrillic # Mn   [2] COMBINING CYRILLIC KAVYKA..COMBINING CYRILLIC PAYEROK</span>
<a id="L115"></a><span class="comment">// See http://www.unicode.org/Public/5.1.0/ucd/UCD.html for full explanation</span>

<a id="L117"></a>type Script struct {
    <a id="L118"></a>lo, hi uint32; <span class="comment">// range of code points</span>
    <a id="L119"></a>script string;
<a id="L120"></a>}

<a id="L122"></a>var chars = make([]Char, MaxChar+1)
<a id="L123"></a>var scripts = make(map[string][]Script)
<a id="L124"></a>var props = make(map[string][]Script) <span class="comment">// a property looks like a script; can share the format</span>

<a id="L126"></a>var lastChar uint32 = 0

<a id="L128"></a><span class="comment">// In UnicodeData.txt, some ranges are marked like this:</span>
<a id="L129"></a><span class="comment">//	3400;&lt;CJK Ideograph Extension A, First&gt;;Lo;0;L;;;;;N;;;;;</span>
<a id="L130"></a><span class="comment">//	4DB5;&lt;CJK Ideograph Extension A, Last&gt;;Lo;0;L;;;;;N;;;;;</span>
<a id="L131"></a><span class="comment">// parseCategory returns a state variable indicating the weirdness.</span>
<a id="L132"></a>type State int

<a id="L134"></a>const (
    <a id="L135"></a>SNormal State = iota; <span class="comment">// known to be zero for the type</span>
    <a id="L136"></a>SFirst;
    <a id="L137"></a>SLast;
    <a id="L138"></a>SMissing;
<a id="L139"></a>)

<a id="L141"></a>func parseCategory(line string) (state State) {
    <a id="L142"></a>field := strings.Split(line, &#34;;&#34;, -1);
    <a id="L143"></a>if len(field) != NumField {
        <a id="L144"></a>die.Logf(&#34;%5s: %d fields (expected %d)\n&#34;, line, len(field), NumField)
    <a id="L145"></a>}
    <a id="L146"></a>point, err := strconv.Btoui64(field[FCodePoint], 16);
    <a id="L147"></a>if err != nil {
        <a id="L148"></a>die.Log(&#34;%.5s...:&#34;, err)
    <a id="L149"></a>}
    <a id="L150"></a>lastChar = uint32(point);
    <a id="L151"></a>if point == 0 {
        <a id="L152"></a>return <span class="comment">// not interesting and we use 0 as unset</span>
    <a id="L153"></a>}
    <a id="L154"></a>if point &gt; MaxChar {
        <a id="L155"></a>return
    <a id="L156"></a>}
    <a id="L157"></a>char := &amp;chars[point];
    <a id="L158"></a>char.field = field;
    <a id="L159"></a>if char.codePoint != 0 {
        <a id="L160"></a>die.Logf(&#34;point U+%04x reused\n&#34;)
    <a id="L161"></a>}
    <a id="L162"></a>char.codePoint = lastChar;
    <a id="L163"></a>char.category = field[FGeneralCategory];
    <a id="L164"></a>category[char.category] = true;
    <a id="L165"></a>switch char.category {
    <a id="L166"></a>case &#34;Nd&#34;:
        <a id="L167"></a><span class="comment">// Decimal digit</span>
        <a id="L168"></a>_, err := strconv.Atoi(field[FNumericValue]);
        <a id="L169"></a>if err != nil {
            <a id="L170"></a>die.Log(&#34;U+%04x: bad numeric field: %s&#34;, point, err)
        <a id="L171"></a>}
    <a id="L172"></a>case &#34;Lu&#34;:
        <a id="L173"></a>char.letter(field[FCodePoint], field[FSimpleLowercaseMapping], field[FSimpleTitlecaseMapping])
    <a id="L174"></a>case &#34;Ll&#34;:
        <a id="L175"></a>char.letter(field[FSimpleUppercaseMapping], field[FCodePoint], field[FSimpleTitlecaseMapping])
    <a id="L176"></a>case &#34;Lt&#34;:
        <a id="L177"></a>char.letter(field[FSimpleUppercaseMapping], field[FSimpleLowercaseMapping], field[FCodePoint])
    <a id="L178"></a>case &#34;Lm&#34;, &#34;Lo&#34;:
        <a id="L179"></a>char.letter(field[FSimpleUppercaseMapping], field[FSimpleLowercaseMapping], field[FSimpleTitlecaseMapping])
    <a id="L180"></a>}
    <a id="L181"></a>switch {
    <a id="L182"></a>case strings.Index(field[FName], &#34;, First&gt;&#34;) &gt; 0:
        <a id="L183"></a>state = SFirst
    <a id="L184"></a>case strings.Index(field[FName], &#34;, Last&gt;&#34;) &gt; 0:
        <a id="L185"></a>state = SLast
    <a id="L186"></a>}
    <a id="L187"></a>return;
<a id="L188"></a>}

<a id="L190"></a>func (char *Char) dump(s string) {
    <a id="L191"></a>fmt.Print(s, &#34; &#34;);
    <a id="L192"></a>for i := 0; i &lt; len(char.field); i++ {
        <a id="L193"></a>fmt.Printf(&#34;%s:%q &#34;, fieldName[i], char.field[i])
    <a id="L194"></a>}
    <a id="L195"></a>fmt.Print(&#34;\n&#34;);
<a id="L196"></a>}

<a id="L198"></a>func (char *Char) letter(u, l, t string) {
    <a id="L199"></a>char.upperCase = char.letterValue(u, &#34;U&#34;);
    <a id="L200"></a>char.lowerCase = char.letterValue(l, &#34;L&#34;);
    <a id="L201"></a>char.titleCase = char.letterValue(t, &#34;T&#34;);
<a id="L202"></a>}

<a id="L204"></a>func (char *Char) letterValue(s string, cas string) int {
    <a id="L205"></a>if s == &#34;&#34; {
        <a id="L206"></a>return 0
    <a id="L207"></a>}
    <a id="L208"></a>v, err := strconv.Btoui64(s, 16);
    <a id="L209"></a>if err != nil {
        <a id="L210"></a>char.dump(cas);
        <a id="L211"></a>die.Logf(&#34;U+%04x: bad letter(%s): %s&#34;, char.codePoint, s, err);
    <a id="L212"></a>}
    <a id="L213"></a>return int(v);
<a id="L214"></a>}

<a id="L216"></a>func allCategories() []string {
    <a id="L217"></a>a := make([]string, len(category));
    <a id="L218"></a>i := 0;
    <a id="L219"></a>for k := range category {
        <a id="L220"></a>a[i] = k;
        <a id="L221"></a>i++;
    <a id="L222"></a>}
    <a id="L223"></a>return a;
<a id="L224"></a>}

<a id="L226"></a>func all(scripts map[string][]Script) []string {
    <a id="L227"></a>a := make([]string, len(scripts));
    <a id="L228"></a>i := 0;
    <a id="L229"></a>for k := range scripts {
        <a id="L230"></a>a[i] = k;
        <a id="L231"></a>i++;
    <a id="L232"></a>}
    <a id="L233"></a>return a;
<a id="L234"></a>}

<a id="L236"></a><span class="comment">// Extract the version number from the URL</span>
<a id="L237"></a>func version() string {
    <a id="L238"></a><span class="comment">// Break on slashes and look for the first numeric field</span>
    <a id="L239"></a>fields := strings.Split(*url, &#34;/&#34;, 0);
    <a id="L240"></a>for _, f := range fields {
        <a id="L241"></a>if len(f) &gt; 0 &amp;&amp; &#39;0&#39; &lt;= f[0] &amp;&amp; f[0] &lt;= &#39;9&#39; {
            <a id="L242"></a>return f
        <a id="L243"></a>}
    <a id="L244"></a>}
    <a id="L245"></a>die.Log(&#34;unknown version&#34;);
    <a id="L246"></a>return &#34;Unknown&#34;;
<a id="L247"></a>}

<a id="L249"></a>func letterOp(code int) bool {
    <a id="L250"></a>switch chars[code].category {
    <a id="L251"></a>case &#34;Lu&#34;, &#34;Ll&#34;, &#34;Lt&#34;, &#34;Lm&#34;, &#34;Lo&#34;:
        <a id="L252"></a>return true
    <a id="L253"></a>}
    <a id="L254"></a>return false;
<a id="L255"></a>}

<a id="L257"></a>func loadChars() {
    <a id="L258"></a>if *dataURL == &#34;&#34; {
        <a id="L259"></a>flag.Set(&#34;data&#34;, *url+&#34;UnicodeData.txt&#34;)
    <a id="L260"></a>}
    <a id="L261"></a>resp, _, err := http.Get(*dataURL);
    <a id="L262"></a>if err != nil {
        <a id="L263"></a>die.Log(err)
    <a id="L264"></a>}
    <a id="L265"></a>if resp.StatusCode != 200 {
        <a id="L266"></a>die.Log(&#34;bad GET status for UnicodeData.txt&#34;, resp.Status)
    <a id="L267"></a>}
    <a id="L268"></a>input := bufio.NewReader(resp.Body);
    <a id="L269"></a>var first uint32 = 0;
    <a id="L270"></a>for {
        <a id="L271"></a>line, err := input.ReadString(&#39;\n&#39;);
        <a id="L272"></a>if err != nil {
            <a id="L273"></a>if err == os.EOF {
                <a id="L274"></a>break
            <a id="L275"></a>}
            <a id="L276"></a>die.Log(err);
        <a id="L277"></a>}
        <a id="L278"></a>switch parseCategory(line[0 : len(line)-1]) {
        <a id="L279"></a>case SNormal:
            <a id="L280"></a>if first != 0 {
                <a id="L281"></a>die.Logf(&#34;bad state normal at U+%04X&#34;, lastChar)
            <a id="L282"></a>}
        <a id="L283"></a>case SFirst:
            <a id="L284"></a>if first != 0 {
                <a id="L285"></a>die.Logf(&#34;bad state first at U+%04X&#34;, lastChar)
            <a id="L286"></a>}
            <a id="L287"></a>first = lastChar;
        <a id="L288"></a>case SLast:
            <a id="L289"></a>if first == 0 {
                <a id="L290"></a>die.Logf(&#34;bad state last at U+%04X&#34;, lastChar)
            <a id="L291"></a>}
            <a id="L292"></a>for i := first + 1; i &lt;= lastChar; i++ {
                <a id="L293"></a>chars[i] = chars[first];
                <a id="L294"></a>chars[i].codePoint = i;
            <a id="L295"></a>}
            <a id="L296"></a>first = 0;
        <a id="L297"></a>}
    <a id="L298"></a>}
    <a id="L299"></a>resp.Body.Close();
<a id="L300"></a>}

<a id="L302"></a>func printCategories() {
    <a id="L303"></a>if *tablelist == &#34;&#34; {
        <a id="L304"></a>return
    <a id="L305"></a>}
    <a id="L306"></a><span class="comment">// Find out which categories to dump</span>
    <a id="L307"></a>list := strings.Split(*tablelist, &#34;,&#34;, 0);
    <a id="L308"></a>if *tablelist == &#34;all&#34; {
        <a id="L309"></a>list = allCategories()
    <a id="L310"></a>}
    <a id="L311"></a>if *test {
        <a id="L312"></a>fullCategoryTest(list);
        <a id="L313"></a>return;
    <a id="L314"></a>}
    <a id="L315"></a>fmt.Printf(
        <a id="L316"></a>&#34;// Generated by running\n&#34;
            <a id="L317"></a>&#34;//	maketables --tables=%s --data=%s\n&#34;
            <a id="L318"></a>&#34;// DO NOT EDIT\n\n&#34;
            <a id="L319"></a>&#34;package unicode\n\n&#34;,
        <a id="L320"></a>*tablelist,
        <a id="L321"></a>*dataURL);

    <a id="L323"></a>fmt.Println(&#34;// Version is the Unicode edition from which the tables are derived.&#34;);
    <a id="L324"></a>fmt.Printf(&#34;const Version = %q\n\n&#34;, version());

    <a id="L326"></a>if *tablelist == &#34;all&#34; {
        <a id="L327"></a>fmt.Println(&#34;// Categories is the set of Unicode data tables.&#34;);
        <a id="L328"></a>fmt.Println(&#34;var Categories = map[string] []Range {&#34;);
        <a id="L329"></a>for k, _ := range category {
            <a id="L330"></a>fmt.Printf(&#34;\t%q: %s,\n&#34;, k, k)
        <a id="L331"></a>}
        <a id="L332"></a>fmt.Printf(&#34;}\n\n&#34;);
    <a id="L333"></a>}

    <a id="L335"></a>decl := make(sort.StringArray, len(list));
    <a id="L336"></a>ndecl := 0;
    <a id="L337"></a>for _, name := range list {
        <a id="L338"></a>if _, ok := category[name]; !ok {
            <a id="L339"></a>die.Log(&#34;unknown category&#34;, name)
        <a id="L340"></a>}
        <a id="L341"></a><span class="comment">// We generate an UpperCase name to serve as concise documentation and an _UnderScored</span>
        <a id="L342"></a><span class="comment">// name to store the data.  This stops godoc dumping all the tables but keeps them</span>
        <a id="L343"></a><span class="comment">// available to clients.</span>
        <a id="L344"></a><span class="comment">// Cases deserving special comments</span>
        <a id="L345"></a>varDecl := &#34;&#34;;
        <a id="L346"></a>switch name {
        <a id="L347"></a>case &#34;letter&#34;:
            <a id="L348"></a>varDecl = &#34;\tLetter = letter;	// Letter is the set of Unicode letters.\n&#34;
        <a id="L349"></a>case &#34;Nd&#34;:
            <a id="L350"></a>varDecl = &#34;\tDigit = _Nd;	// Digit is the set of Unicode characters with the \&#34;decimal digit\&#34; property.\n&#34;
        <a id="L351"></a>case &#34;Lu&#34;:
            <a id="L352"></a>varDecl = &#34;\tUpper = _Lu;	// Upper is the set of Unicode upper case letters.\n&#34;
        <a id="L353"></a>case &#34;Ll&#34;:
            <a id="L354"></a>varDecl = &#34;\tLower = _Ll;	// Lower is the set of Unicode lower case letters.\n&#34;
        <a id="L355"></a>case &#34;Lt&#34;:
            <a id="L356"></a>varDecl = &#34;\tTitle = _Lt;	// Title is the set of Unicode title case letters.\n&#34;
        <a id="L357"></a>}
        <a id="L358"></a>if name != &#34;letter&#34; {
            <a id="L359"></a>varDecl += fmt.Sprintf(
                <a id="L360"></a>&#34;\t%s = _%s;	// %s is the set of Unicode characters in category %s.\n&#34;,
                <a id="L361"></a>name, name, name, name)
        <a id="L362"></a>}
        <a id="L363"></a>decl[ndecl] = varDecl;
        <a id="L364"></a>ndecl++;
        <a id="L365"></a>if name == &#34;letter&#34; { <span class="comment">// special case</span>
            <a id="L366"></a>dumpRange(
                <a id="L367"></a>&#34;var letter = []Range {\n&#34;,
                <a id="L368"></a>letterOp);
            <a id="L369"></a>continue;
        <a id="L370"></a>}
        <a id="L371"></a>dumpRange(
            <a id="L372"></a>fmt.Sprintf(&#34;var _%s = []Range {\n&#34;, name),
            <a id="L373"></a>func(code int) bool { return chars[code].category == name });
    <a id="L374"></a>}
    <a id="L375"></a>decl.Sort();
    <a id="L376"></a>fmt.Println(&#34;var (&#34;);
    <a id="L377"></a>for _, d := range decl {
        <a id="L378"></a>fmt.Print(d)
    <a id="L379"></a>}
    <a id="L380"></a>fmt.Println(&#34;)\n&#34;);
<a id="L381"></a>}

<a id="L383"></a>type Op func(code int) bool

<a id="L385"></a>const format = &#34;\tRange{0x%04x, 0x%04x, %d},\n&#34;

<a id="L387"></a>func dumpRange(header string, inCategory Op) {
    <a id="L388"></a>fmt.Print(header);
    <a id="L389"></a>next := 0;
    <a id="L390"></a><span class="comment">// one Range for each iteration</span>
    <a id="L391"></a>for {
        <a id="L392"></a><span class="comment">// look for start of range</span>
        <a id="L393"></a>for next &lt; len(chars) &amp;&amp; !inCategory(next) {
            <a id="L394"></a>next++
        <a id="L395"></a>}
        <a id="L396"></a>if next &gt;= len(chars) {
            <a id="L397"></a><span class="comment">// no characters remain</span>
            <a id="L398"></a>break
        <a id="L399"></a>}

        <a id="L401"></a><span class="comment">// start of range</span>
        <a id="L402"></a>lo := next;
        <a id="L403"></a>hi := next;
        <a id="L404"></a>stride := 1;
        <a id="L405"></a><span class="comment">// accept lo</span>
        <a id="L406"></a>next++;
        <a id="L407"></a><span class="comment">// look for another character to set the stride</span>
        <a id="L408"></a>for next &lt; len(chars) &amp;&amp; !inCategory(next) {
            <a id="L409"></a>next++
        <a id="L410"></a>}
        <a id="L411"></a>if next &gt;= len(chars) {
            <a id="L412"></a><span class="comment">// no more characters</span>
            <a id="L413"></a>fmt.Printf(format, lo, hi, stride);
            <a id="L414"></a>break;
        <a id="L415"></a>}
        <a id="L416"></a><span class="comment">// set stride</span>
        <a id="L417"></a>stride = next - lo;
        <a id="L418"></a><span class="comment">// check for length of run. next points to first jump in stride</span>
        <a id="L419"></a>for i := next; i &lt; len(chars); i++ {
            <a id="L420"></a>if inCategory(i) == (((i - lo) % stride) == 0) {
                <a id="L421"></a><span class="comment">// accept</span>
                <a id="L422"></a>if inCategory(i) {
                    <a id="L423"></a>hi = i
                <a id="L424"></a>}
            <a id="L425"></a>} else {
                <a id="L426"></a><span class="comment">// no more characters in this run</span>
                <a id="L427"></a>break
            <a id="L428"></a>}
        <a id="L429"></a>}
        <a id="L430"></a>fmt.Printf(format, lo, hi, stride);
        <a id="L431"></a><span class="comment">// next range: start looking where this range ends</span>
        <a id="L432"></a>next = hi + 1;
    <a id="L433"></a>}
    <a id="L434"></a>fmt.Print(&#34;}\n\n&#34;);
<a id="L435"></a>}

<a id="L437"></a>func fullCategoryTest(list []string) {
    <a id="L438"></a>for _, name := range list {
        <a id="L439"></a>if _, ok := category[name]; !ok {
            <a id="L440"></a>die.Log(&#34;unknown category&#34;, name)
        <a id="L441"></a>}
        <a id="L442"></a>r, ok := unicode.Categories[name];
        <a id="L443"></a>if !ok {
            <a id="L444"></a>die.Log(&#34;unknown table&#34;, name)
        <a id="L445"></a>}
        <a id="L446"></a>if name == &#34;letter&#34; {
            <a id="L447"></a>verifyRange(name, letterOp, r)
        <a id="L448"></a>} else {
            <a id="L449"></a>verifyRange(
                <a id="L450"></a>name,
                <a id="L451"></a>func(code int) bool { return chars[code].category == name },
                <a id="L452"></a>r)
        <a id="L453"></a>}
    <a id="L454"></a>}
<a id="L455"></a>}

<a id="L457"></a>func verifyRange(name string, inCategory Op, table []unicode.Range) {
    <a id="L458"></a>for i := range chars {
        <a id="L459"></a>web := inCategory(i);
        <a id="L460"></a>pkg := unicode.Is(table, i);
        <a id="L461"></a>if web != pkg {
            <a id="L462"></a>fmt.Fprintf(os.Stderr, &#34;%s: U+%04X: web=%t pkg=%t\n&#34;, name, i, web, pkg)
        <a id="L463"></a>}
    <a id="L464"></a>}
<a id="L465"></a>}

<a id="L467"></a>func parseScript(line string, scripts map[string][]Script) {
    <a id="L468"></a>comment := strings.Index(line, &#34;#&#34;);
    <a id="L469"></a>if comment &gt;= 0 {
        <a id="L470"></a>line = line[0:comment]
    <a id="L471"></a>}
    <a id="L472"></a>line = strings.TrimSpace(line);
    <a id="L473"></a>if len(line) == 0 {
        <a id="L474"></a>return
    <a id="L475"></a>}
    <a id="L476"></a>field := strings.Split(line, &#34;;&#34;, -1);
    <a id="L477"></a>if len(field) != 2 {
        <a id="L478"></a>die.Logf(&#34;%s: %d fields (expected 2)\n&#34;, line, len(field))
    <a id="L479"></a>}
    <a id="L480"></a>matches := scriptRe.MatchStrings(line);
    <a id="L481"></a>if len(matches) != 4 {
        <a id="L482"></a>die.Logf(&#34;%s: %d matches (expected 3)\n&#34;, line, len(matches))
    <a id="L483"></a>}
    <a id="L484"></a>lo, err := strconv.Btoui64(matches[1], 16);
    <a id="L485"></a>if err != nil {
        <a id="L486"></a>die.Log(&#34;%.5s...:&#34;, err)
    <a id="L487"></a>}
    <a id="L488"></a>hi := lo;
    <a id="L489"></a>if len(matches[2]) &gt; 2 { <span class="comment">// ignore leading ..</span>
        <a id="L490"></a>hi, err = strconv.Btoui64(matches[2][2:len(matches[2])], 16);
        <a id="L491"></a>if err != nil {
            <a id="L492"></a>die.Log(&#34;%.5s...:&#34;, err)
        <a id="L493"></a>}
    <a id="L494"></a>}
    <a id="L495"></a>name := matches[3];
    <a id="L496"></a>s, ok := scripts[name];
    <a id="L497"></a>if !ok || len(s) == cap(s) {
        <a id="L498"></a>ns := make([]Script, len(s), len(s)+100);
        <a id="L499"></a>for i, sc := range s {
            <a id="L500"></a>ns[i] = sc
        <a id="L501"></a>}
        <a id="L502"></a>s = ns;
    <a id="L503"></a>}
    <a id="L504"></a>s = s[0 : len(s)+1];
    <a id="L505"></a>s[len(s)-1] = Script{uint32(lo), uint32(hi), name};
    <a id="L506"></a>scripts[name] = s;
<a id="L507"></a>}

<a id="L509"></a><span class="comment">// The script tables have a lot of adjacent elements. Fold them together.</span>
<a id="L510"></a>func foldAdjacent(r []Script) []unicode.Range {
    <a id="L511"></a>s := make([]unicode.Range, 0, len(r));
    <a id="L512"></a>j := 0;
    <a id="L513"></a>for i := 0; i &lt; len(r); i++ {
        <a id="L514"></a>if j &gt; 0 &amp;&amp; int(r[i].lo) == s[j-1].Hi+1 {
            <a id="L515"></a>s[j-1].Hi = int(r[i].hi)
        <a id="L516"></a>} else {
            <a id="L517"></a>s = s[0 : j+1];
            <a id="L518"></a>s[j] = unicode.Range{int(r[i].lo), int(r[i].hi), 1};
            <a id="L519"></a>j++;
        <a id="L520"></a>}
    <a id="L521"></a>}
    <a id="L522"></a>return s;
<a id="L523"></a>}

<a id="L525"></a>func fullScriptTest(list []string, installed map[string][]unicode.Range, scripts map[string][]Script) {
    <a id="L526"></a>for _, name := range list {
        <a id="L527"></a>if _, ok := scripts[name]; !ok {
            <a id="L528"></a>die.Log(&#34;unknown script&#34;, name)
        <a id="L529"></a>}
        <a id="L530"></a>_, ok := installed[name];
        <a id="L531"></a>if !ok {
            <a id="L532"></a>die.Log(&#34;unknown table&#34;, name)
        <a id="L533"></a>}
        <a id="L534"></a>for _, script := range scripts[name] {
            <a id="L535"></a>for r := script.lo; r &lt;= script.hi; r++ {
                <a id="L536"></a>if !unicode.Is(installed[name], int(r)) {
                    <a id="L537"></a>fmt.Fprintf(os.Stderr, &#34;U+%04X: not in script %s\n&#34;, r, name)
                <a id="L538"></a>}
            <a id="L539"></a>}
        <a id="L540"></a>}
    <a id="L541"></a>}
<a id="L542"></a>}

<a id="L544"></a><span class="comment">// PropList.txt has the same format as Scripts.txt so we can share its parser.</span>
<a id="L545"></a>func printScriptOrProperty(doProps bool) {
    <a id="L546"></a>flag := &#34;scripts&#34;;
    <a id="L547"></a>flaglist := *scriptlist;
    <a id="L548"></a>file := &#34;Scripts.txt&#34;;
    <a id="L549"></a>table := scripts;
    <a id="L550"></a>installed := unicode.Scripts;
    <a id="L551"></a>if doProps {
        <a id="L552"></a>flag = &#34;props&#34;;
        <a id="L553"></a>flaglist = *proplist;
        <a id="L554"></a>file = &#34;PropList.txt&#34;;
        <a id="L555"></a>table = props;
        <a id="L556"></a>installed = unicode.Properties;
    <a id="L557"></a>}
    <a id="L558"></a>if flaglist == &#34;&#34; {
        <a id="L559"></a>return
    <a id="L560"></a>}
    <a id="L561"></a>var err os.Error;
    <a id="L562"></a>resp, _, err := http.Get(*url + file);
    <a id="L563"></a>if err != nil {
        <a id="L564"></a>die.Log(err)
    <a id="L565"></a>}
    <a id="L566"></a>if resp.StatusCode != 200 {
        <a id="L567"></a>die.Log(&#34;bad GET status for &#34;, file, &#34;:&#34;, resp.Status)
    <a id="L568"></a>}
    <a id="L569"></a>input := bufio.NewReader(resp.Body);
    <a id="L570"></a>for {
        <a id="L571"></a>line, err := input.ReadString(&#39;\n&#39;);
        <a id="L572"></a>if err != nil {
            <a id="L573"></a>if err == os.EOF {
                <a id="L574"></a>break
            <a id="L575"></a>}
            <a id="L576"></a>die.Log(err);
        <a id="L577"></a>}
        <a id="L578"></a>parseScript(line[0:len(line)-1], table);
    <a id="L579"></a>}
    <a id="L580"></a>resp.Body.Close();

    <a id="L582"></a><span class="comment">// Find out which scripts to dump</span>
    <a id="L583"></a>list := strings.Split(flaglist, &#34;,&#34;, 0);
    <a id="L584"></a>if flaglist == &#34;all&#34; {
        <a id="L585"></a>list = all(table)
    <a id="L586"></a>}
    <a id="L587"></a>if *test {
        <a id="L588"></a>fullScriptTest(list, installed, table);
        <a id="L589"></a>return;
    <a id="L590"></a>}

    <a id="L592"></a>fmt.Printf(
        <a id="L593"></a>&#34;// Generated by running\n&#34;
            <a id="L594"></a>&#34;//	maketables --%s=%s --url=%s\n&#34;
            <a id="L595"></a>&#34;// DO NOT EDIT\n\n&#34;,
        <a id="L596"></a>flag,
        <a id="L597"></a>flaglist,
        <a id="L598"></a>*url);
    <a id="L599"></a>if flaglist == &#34;all&#34; {
        <a id="L600"></a>if doProps {
            <a id="L601"></a>fmt.Println(&#34;// Properties is the set of Unicode property tables.&#34;);
            <a id="L602"></a>fmt.Println(&#34;var Properties = map[string] []Range {&#34;);
        <a id="L603"></a>} else {
            <a id="L604"></a>fmt.Println(&#34;// Scripts is the set of Unicode script tables.&#34;);
            <a id="L605"></a>fmt.Println(&#34;var Scripts = map[string] []Range {&#34;);
        <a id="L606"></a>}
        <a id="L607"></a>for k, _ := range table {
            <a id="L608"></a>fmt.Printf(&#34;\t%q: %s,\n&#34;, k, k)
        <a id="L609"></a>}
        <a id="L610"></a>fmt.Printf(&#34;}\n\n&#34;);
    <a id="L611"></a>}

    <a id="L613"></a>decl := make(sort.StringArray, len(list));
    <a id="L614"></a>ndecl := 0;
    <a id="L615"></a>for _, name := range list {
        <a id="L616"></a>if doProps {
            <a id="L617"></a>decl[ndecl] = fmt.Sprintf(
                <a id="L618"></a>&#34;\t%s = _%s;\t// %s is the set of Unicode characters with property %s.\n&#34;,
                <a id="L619"></a>name, name, name, name)
        <a id="L620"></a>} else {
            <a id="L621"></a>decl[ndecl] = fmt.Sprintf(
                <a id="L622"></a>&#34;\t%s = _%s;\t// %s is the set of Unicode characters in script %s.\n&#34;,
                <a id="L623"></a>name, name, name, name)
        <a id="L624"></a>}
        <a id="L625"></a>ndecl++;
        <a id="L626"></a>fmt.Printf(&#34;var _%s = []Range {\n&#34;, name);
        <a id="L627"></a>ranges := foldAdjacent(table[name]);
        <a id="L628"></a>for _, s := range ranges {
            <a id="L629"></a>fmt.Printf(format, s.Lo, s.Hi, s.Stride)
        <a id="L630"></a>}
        <a id="L631"></a>fmt.Printf(&#34;}\n\n&#34;);
    <a id="L632"></a>}
    <a id="L633"></a>decl.Sort();
    <a id="L634"></a>fmt.Println(&#34;var (&#34;);
    <a id="L635"></a>for _, d := range decl {
        <a id="L636"></a>fmt.Print(d)
    <a id="L637"></a>}
    <a id="L638"></a>fmt.Println(&#34;)\n&#34;);
<a id="L639"></a>}

<a id="L641"></a>const (
    <a id="L642"></a>CaseUpper = 1 &lt;&lt; iota;
    <a id="L643"></a>CaseLower;
    <a id="L644"></a>CaseTitle;
    <a id="L645"></a>CaseNone    = 0;  <span class="comment">// must be zero</span>
    <a id="L646"></a>CaseMissing = -1; <span class="comment">// character not present; not a valid case state</span>
<a id="L647"></a>)

<a id="L649"></a>type caseState struct {
    <a id="L650"></a>point        int;
    <a id="L651"></a>_case        int;
    <a id="L652"></a>deltaToUpper int;
    <a id="L653"></a>deltaToLower int;
    <a id="L654"></a>deltaToTitle int;
<a id="L655"></a>}

<a id="L657"></a><span class="comment">// Is d a continuation of the state of c?</span>
<a id="L658"></a>func (c *caseState) adjacent(d *caseState) bool {
    <a id="L659"></a>if d.point &lt; c.point {
        <a id="L660"></a>c, d = d, c
    <a id="L661"></a>}
    <a id="L662"></a>switch {
    <a id="L663"></a>case d.point != c.point+1: <span class="comment">// code points not adjacent (shouldn&#39;t happen)</span>
        <a id="L664"></a>return false
    <a id="L665"></a>case d._case != c._case: <span class="comment">// different cases</span>
        <a id="L666"></a>return c.upperLowerAdjacent(d)
    <a id="L667"></a>case c._case == CaseNone:
        <a id="L668"></a>return false
    <a id="L669"></a>case c._case == CaseMissing:
        <a id="L670"></a>return false
    <a id="L671"></a>case d.deltaToUpper != c.deltaToUpper:
        <a id="L672"></a>return false
    <a id="L673"></a>case d.deltaToLower != c.deltaToLower:
        <a id="L674"></a>return false
    <a id="L675"></a>case d.deltaToTitle != c.deltaToTitle:
        <a id="L676"></a>return false
    <a id="L677"></a>}
    <a id="L678"></a>return true;
<a id="L679"></a>}

<a id="L681"></a><span class="comment">// Is d the same as c, but opposite in upper/lower case? this would make it</span>
<a id="L682"></a><span class="comment">// an element of an UpperLower sequence.</span>
<a id="L683"></a>func (c *caseState) upperLowerAdjacent(d *caseState) bool {
    <a id="L684"></a><span class="comment">// check they&#39;re a matched case pair.  we know they have adjacent values</span>
    <a id="L685"></a>switch {
    <a id="L686"></a>case c._case == CaseUpper &amp;&amp; d._case != CaseLower:
        <a id="L687"></a>return false
    <a id="L688"></a>case c._case == CaseLower &amp;&amp; d._case != CaseUpper:
        <a id="L689"></a>return false
    <a id="L690"></a>}
    <a id="L691"></a><span class="comment">// matched pair (at least in upper/lower).  make the order Upper Lower</span>
    <a id="L692"></a>if c._case == CaseLower {
        <a id="L693"></a>c, d = d, c
    <a id="L694"></a>}
    <a id="L695"></a><span class="comment">// for an Upper Lower sequence the deltas have to be in order</span>
    <a id="L696"></a><span class="comment">//	c: 0 1 0</span>
    <a id="L697"></a><span class="comment">//	d: -1 0 -1</span>
    <a id="L698"></a>switch {
    <a id="L699"></a>case c.deltaToUpper != 0:
        <a id="L700"></a>return false
    <a id="L701"></a>case c.deltaToLower != 1:
        <a id="L702"></a>return false
    <a id="L703"></a>case c.deltaToTitle != 0:
        <a id="L704"></a>return false
    <a id="L705"></a>case d.deltaToUpper != -1:
        <a id="L706"></a>return false
    <a id="L707"></a>case d.deltaToLower != 0:
        <a id="L708"></a>return false
    <a id="L709"></a>case d.deltaToTitle != -1:
        <a id="L710"></a>return false
    <a id="L711"></a>}
    <a id="L712"></a>return true;
<a id="L713"></a>}

<a id="L715"></a><span class="comment">// Does this character start an UpperLower sequence?</span>
<a id="L716"></a>func (c *caseState) isUpperLower() bool {
    <a id="L717"></a><span class="comment">// for an Upper Lower sequence the deltas have to be in order</span>
    <a id="L718"></a><span class="comment">//	c: 0 1 0</span>
    <a id="L719"></a>switch {
    <a id="L720"></a>case c.deltaToUpper != 0:
        <a id="L721"></a>return false
    <a id="L722"></a>case c.deltaToLower != 1:
        <a id="L723"></a>return false
    <a id="L724"></a>case c.deltaToTitle != 0:
        <a id="L725"></a>return false
    <a id="L726"></a>}
    <a id="L727"></a>return true;
<a id="L728"></a>}

<a id="L730"></a><span class="comment">// Does this character start a LowerUpper sequence?</span>
<a id="L731"></a>func (c *caseState) isLowerUpper() bool {
    <a id="L732"></a><span class="comment">// for an Upper Lower sequence the deltas have to be in order</span>
    <a id="L733"></a><span class="comment">//	c: -1 0 -1</span>
    <a id="L734"></a>switch {
    <a id="L735"></a>case c.deltaToUpper != -1:
        <a id="L736"></a>return false
    <a id="L737"></a>case c.deltaToLower != 0:
        <a id="L738"></a>return false
    <a id="L739"></a>case c.deltaToTitle != -1:
        <a id="L740"></a>return false
    <a id="L741"></a>}
    <a id="L742"></a>return true;
<a id="L743"></a>}

<a id="L745"></a>func getCaseState(i int) (c *caseState) {
    <a id="L746"></a>c = &amp;caseState{point: i, _case: CaseNone};
    <a id="L747"></a>ch := &amp;chars[i];
    <a id="L748"></a>switch int(ch.codePoint) {
    <a id="L749"></a>case 0:
        <a id="L750"></a>c._case = CaseMissing; <span class="comment">// Will get NUL wrong but that doesn&#39;t matter</span>
        <a id="L751"></a>return;
    <a id="L752"></a>case ch.upperCase:
        <a id="L753"></a>c._case = CaseUpper
    <a id="L754"></a>case ch.lowerCase:
        <a id="L755"></a>c._case = CaseLower
    <a id="L756"></a>case ch.titleCase:
        <a id="L757"></a>c._case = CaseTitle
    <a id="L758"></a>}
    <a id="L759"></a>if ch.upperCase != 0 {
        <a id="L760"></a>c.deltaToUpper = ch.upperCase - i
    <a id="L761"></a>}
    <a id="L762"></a>if ch.lowerCase != 0 {
        <a id="L763"></a>c.deltaToLower = ch.lowerCase - i
    <a id="L764"></a>}
    <a id="L765"></a>if ch.titleCase != 0 {
        <a id="L766"></a>c.deltaToTitle = ch.titleCase - i
    <a id="L767"></a>}
    <a id="L768"></a>return;
<a id="L769"></a>}

<a id="L771"></a>func printCases() {
    <a id="L772"></a>if !*cases {
        <a id="L773"></a>return
    <a id="L774"></a>}
    <a id="L775"></a>if *test {
        <a id="L776"></a>fullCaseTest();
        <a id="L777"></a>return;
    <a id="L778"></a>}
    <a id="L779"></a>fmt.Printf(
        <a id="L780"></a>&#34;// Generated by running\n&#34;
            <a id="L781"></a>&#34;//	maketables --data=%s\n&#34;
            <a id="L782"></a>&#34;// DO NOT EDIT\n\n&#34;
            <a id="L783"></a>&#34;// CaseRanges is the table describing case mappings for all letters with\n&#34;
            <a id="L784"></a>&#34;// non-self mappings.\n&#34;
            <a id="L785"></a>&#34;var CaseRanges = _CaseRanges\n&#34;
            <a id="L786"></a>&#34;var _CaseRanges = []CaseRange {\n&#34;,
        <a id="L787"></a>*dataURL);

    <a id="L789"></a>var startState *caseState;    <span class="comment">// the start of a run; nil for not active</span>
    <a id="L790"></a>var prevState = &amp;caseState{}; <span class="comment">// the state of the previous character</span>
    <a id="L791"></a>for i := range chars {
        <a id="L792"></a>state := getCaseState(i);
        <a id="L793"></a>if state.adjacent(prevState) {
            <a id="L794"></a>prevState = state;
            <a id="L795"></a>continue;
        <a id="L796"></a>}
        <a id="L797"></a><span class="comment">// end of run (possibly)</span>
        <a id="L798"></a>printCaseRange(startState, prevState);
        <a id="L799"></a>startState = nil;
        <a id="L800"></a>if state._case != CaseMissing &amp;&amp; state._case != CaseNone {
            <a id="L801"></a>startState = state
        <a id="L802"></a>}
        <a id="L803"></a>prevState = state;
    <a id="L804"></a>}
    <a id="L805"></a>fmt.Printf(&#34;}\n&#34;);
<a id="L806"></a>}

<a id="L808"></a>func printCaseRange(lo, hi *caseState) {
    <a id="L809"></a>if lo == nil {
        <a id="L810"></a>return
    <a id="L811"></a>}
    <a id="L812"></a>if lo.deltaToUpper == 0 &amp;&amp; lo.deltaToLower == 0 &amp;&amp; lo.deltaToTitle == 0 {
        <a id="L813"></a><span class="comment">// character represents itself in all cases - no need to mention it</span>
        <a id="L814"></a>return
    <a id="L815"></a>}
    <a id="L816"></a>switch {
    <a id="L817"></a>case hi.point &gt; lo.point &amp;&amp; lo.isUpperLower():
        <a id="L818"></a>fmt.Printf(&#34;\tCaseRange{0x%04X, 0x%04X, d{UpperLower, UpperLower, UpperLower}},\n&#34;,
            <a id="L819"></a>lo.point, hi.point)
    <a id="L820"></a>case hi.point &gt; lo.point &amp;&amp; lo.isLowerUpper():
        <a id="L821"></a>die.Log(&#34;LowerUpper sequence: should not happen: U+%04X.  If it&#39;s real, need to fix To()&#34;, lo.point);
        <a id="L822"></a>fmt.Printf(&#34;\tCaseRange{0x%04X, 0x%04X, d{LowerUpper, LowerUpper, LowerUpper}},\n&#34;,
            <a id="L823"></a>lo.point, hi.point);
    <a id="L824"></a>default:
        <a id="L825"></a>fmt.Printf(&#34;\tCaseRange{0x%04X, 0x%04X, d{%d, %d, %d}},\n&#34;,
            <a id="L826"></a>lo.point, hi.point,
            <a id="L827"></a>lo.deltaToUpper, lo.deltaToLower, lo.deltaToTitle)
    <a id="L828"></a>}
<a id="L829"></a>}

<a id="L831"></a><span class="comment">// If the cased value in the Char is 0, it means use the rune itself.</span>
<a id="L832"></a>func caseIt(rune, cased int) int {
    <a id="L833"></a>if cased == 0 {
        <a id="L834"></a>return rune
    <a id="L835"></a>}
    <a id="L836"></a>return cased;
<a id="L837"></a>}

<a id="L839"></a>func fullCaseTest() {
    <a id="L840"></a>for i, c := range chars {
        <a id="L841"></a>lower := unicode.ToLower(i);
        <a id="L842"></a>want := caseIt(i, c.lowerCase);
        <a id="L843"></a>if lower != want {
            <a id="L844"></a>fmt.Fprintf(os.Stderr, &#34;lower U+%04X should be U+%04X is U+%04X\n&#34;, i, want, lower)
        <a id="L845"></a>}
        <a id="L846"></a>upper := unicode.ToUpper(i);
        <a id="L847"></a>want = caseIt(i, c.upperCase);
        <a id="L848"></a>if upper != want {
            <a id="L849"></a>fmt.Fprintf(os.Stderr, &#34;upper U+%04X should be U+%04X is U+%04X\n&#34;, i, want, upper)
        <a id="L850"></a>}
        <a id="L851"></a>title := unicode.ToTitle(i);
        <a id="L852"></a>want = caseIt(i, c.titleCase);
        <a id="L853"></a>if title != want {
            <a id="L854"></a>fmt.Fprintf(os.Stderr, &#34;title U+%04X should be U+%04X is U+%04X\n&#34;, i, want, title)
        <a id="L855"></a>}
    <a id="L856"></a>}
<a id="L857"></a>}
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
