<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/nacl/av/event.go</title>

  <link rel="stylesheet" type="text/css" href="../../../../../doc/style.css">
  <script type="text/javascript" src="../../../../../doc/godocs.js"></script>

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
        <a href="../../../../../index.html"><img src="../../../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../../../index.html">Source files</a></li>

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
  <h1 id="generatedHeader">Source file /src/pkg/exp/nacl/av/event.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// NaCl GUI events.</span>
<a id="L6"></a><span class="comment">// Clients do not have raw access to the event stream</span>
<a id="L7"></a><span class="comment">// (only filtered through the lens of package draw)</span>
<a id="L8"></a><span class="comment">// but perhaps they will.</span>

<a id="L10"></a>package av

<a id="L12"></a>import (
    <a id="L13"></a>&#34;bytes&#34;;
    <a id="L14"></a>&#34;debug/binary&#34;;
    <a id="L15"></a>&#34;exp/draw&#34;;
    <a id="L16"></a>&#34;log&#34;;
    <a id="L17"></a>&#34;os&#34;;
    <a id="L18"></a>&#34;time&#34;;
<a id="L19"></a>)

<a id="L21"></a><span class="comment">// An eventType identifies the type of a Native Client Event.</span>
<a id="L22"></a>type eventType uint8

<a id="L24"></a>const (
    <a id="L25"></a>eventActive = 1 + iota;
    <a id="L26"></a>eventExpose;
    <a id="L27"></a>eventKeyDown;
    <a id="L28"></a>eventKeyUp;
    <a id="L29"></a>eventMouseMotion;
    <a id="L30"></a>eventMouseButtonDown;
    <a id="L31"></a>eventMouseButtonUp;
    <a id="L32"></a>eventQuit;
    <a id="L33"></a>eventUnsupported;
<a id="L34"></a>)

<a id="L36"></a><span class="comment">// A key represents a key on a keyboard.</span>
<a id="L37"></a>type key uint16

<a id="L39"></a>const (
    <a id="L40"></a>keyUnknown      = 0;
    <a id="L41"></a>keyFirst        = 0;
    <a id="L42"></a>keyBackspace    = 8;
    <a id="L43"></a>keyTab          = 9;
    <a id="L44"></a>keyClear        = 12;
    <a id="L45"></a>keyReturn       = 13;
    <a id="L46"></a>keyPause        = 19;
    <a id="L47"></a>keyEscape       = 27;
    <a id="L48"></a>keySpace        = 32;
    <a id="L49"></a>keyExclaim      = 33;
    <a id="L50"></a>keyQuotedbl     = 34;
    <a id="L51"></a>keyHash         = 35;
    <a id="L52"></a>keyDollar       = 36;
    <a id="L53"></a>keyAmpersand    = 38;
    <a id="L54"></a>keyQuote        = 39;
    <a id="L55"></a>keyLeftparen    = 40;
    <a id="L56"></a>keyRightparen   = 41;
    <a id="L57"></a>keyAsterisk     = 42;
    <a id="L58"></a>keyPlus         = 43;
    <a id="L59"></a>keyComma        = 44;
    <a id="L60"></a>keyMinus        = 45;
    <a id="L61"></a>keyPeriod       = 46;
    <a id="L62"></a>keySlash        = 47;
    <a id="L63"></a>key0            = 48;
    <a id="L64"></a>key1            = 49;
    <a id="L65"></a>key2            = 50;
    <a id="L66"></a>key3            = 51;
    <a id="L67"></a>key4            = 52;
    <a id="L68"></a>key5            = 53;
    <a id="L69"></a>key6            = 54;
    <a id="L70"></a>key7            = 55;
    <a id="L71"></a>key8            = 56;
    <a id="L72"></a>key9            = 57;
    <a id="L73"></a>keyColon        = 58;
    <a id="L74"></a>keySemicolon    = 59;
    <a id="L75"></a>keyLess         = 60;
    <a id="L76"></a>keyEquals       = 61;
    <a id="L77"></a>keyGreater      = 62;
    <a id="L78"></a>keyQuestion     = 63;
    <a id="L79"></a>keyAt           = 64;
    <a id="L80"></a>keyLeftbracket  = 91;
    <a id="L81"></a>keyBackslash    = 92;
    <a id="L82"></a>keyRightbracket = 93;
    <a id="L83"></a>keyCaret        = 94;
    <a id="L84"></a>keyUnderscore   = 95;
    <a id="L85"></a>keyBackquote    = 96;
    <a id="L86"></a>keyA            = 97;
    <a id="L87"></a>keyB            = 98;
    <a id="L88"></a>keyC            = 99;
    <a id="L89"></a>keyD            = 100;
    <a id="L90"></a>keyE            = 101;
    <a id="L91"></a>keyF            = 102;
    <a id="L92"></a>keyG            = 103;
    <a id="L93"></a>keyH            = 104;
    <a id="L94"></a>keyI            = 105;
    <a id="L95"></a>keyJ            = 106;
    <a id="L96"></a>keyK            = 107;
    <a id="L97"></a>keyL            = 108;
    <a id="L98"></a>keyM            = 109;
    <a id="L99"></a>keyN            = 110;
    <a id="L100"></a>keyO            = 111;
    <a id="L101"></a>keyP            = 112;
    <a id="L102"></a>keyQ            = 113;
    <a id="L103"></a>keyR            = 114;
    <a id="L104"></a>keyS            = 115;
    <a id="L105"></a>keyT            = 116;
    <a id="L106"></a>keyU            = 117;
    <a id="L107"></a>keyV            = 118;
    <a id="L108"></a>keyW            = 119;
    <a id="L109"></a>keyX            = 120;
    <a id="L110"></a>keyY            = 121;
    <a id="L111"></a>keyZ            = 122;
    <a id="L112"></a>keyDelete       = 127;
    <a id="L113"></a>keyWorld0       = 160;
    <a id="L114"></a>keyWorld1       = 161;
    <a id="L115"></a>keyWorld2       = 162;
    <a id="L116"></a>keyWorld3       = 163;
    <a id="L117"></a>keyWorld4       = 164;
    <a id="L118"></a>keyWorld5       = 165;
    <a id="L119"></a>keyWorld6       = 166;
    <a id="L120"></a>keyWorld7       = 167;
    <a id="L121"></a>keyWorld8       = 168;
    <a id="L122"></a>keyWorld9       = 169;
    <a id="L123"></a>keyWorld10      = 170;
    <a id="L124"></a>keyWorld11      = 171;
    <a id="L125"></a>keyWorld12      = 172;
    <a id="L126"></a>keyWorld13      = 173;
    <a id="L127"></a>keyWorld14      = 174;
    <a id="L128"></a>keyWorld15      = 175;
    <a id="L129"></a>keyWorld16      = 176;
    <a id="L130"></a>keyWorld17      = 177;
    <a id="L131"></a>keyWorld18      = 178;
    <a id="L132"></a>keyWorld19      = 179;
    <a id="L133"></a>keyWorld20      = 180;
    <a id="L134"></a>keyWorld21      = 181;
    <a id="L135"></a>keyWorld22      = 182;
    <a id="L136"></a>keyWorld23      = 183;
    <a id="L137"></a>keyWorld24      = 184;
    <a id="L138"></a>keyWorld25      = 185;
    <a id="L139"></a>keyWorld26      = 186;
    <a id="L140"></a>keyWorld27      = 187;
    <a id="L141"></a>keyWorld28      = 188;
    <a id="L142"></a>keyWorld29      = 189;
    <a id="L143"></a>keyWorld30      = 190;
    <a id="L144"></a>keyWorld31      = 191;
    <a id="L145"></a>keyWorld32      = 192;
    <a id="L146"></a>keyWorld33      = 193;
    <a id="L147"></a>keyWorld34      = 194;
    <a id="L148"></a>keyWorld35      = 195;
    <a id="L149"></a>keyWorld36      = 196;
    <a id="L150"></a>keyWorld37      = 197;
    <a id="L151"></a>keyWorld38      = 198;
    <a id="L152"></a>keyWorld39      = 199;
    <a id="L153"></a>keyWorld40      = 200;
    <a id="L154"></a>keyWorld41      = 201;
    <a id="L155"></a>keyWorld42      = 202;
    <a id="L156"></a>keyWorld43      = 203;
    <a id="L157"></a>keyWorld44      = 204;
    <a id="L158"></a>keyWorld45      = 205;
    <a id="L159"></a>keyWorld46      = 206;
    <a id="L160"></a>keyWorld47      = 207;
    <a id="L161"></a>keyWorld48      = 208;
    <a id="L162"></a>keyWorld49      = 209;
    <a id="L163"></a>keyWorld50      = 210;
    <a id="L164"></a>keyWorld51      = 211;
    <a id="L165"></a>keyWorld52      = 212;
    <a id="L166"></a>keyWorld53      = 213;
    <a id="L167"></a>keyWorld54      = 214;
    <a id="L168"></a>keyWorld55      = 215;
    <a id="L169"></a>keyWorld56      = 216;
    <a id="L170"></a>keyWorld57      = 217;
    <a id="L171"></a>keyWorld58      = 218;
    <a id="L172"></a>keyWorld59      = 219;
    <a id="L173"></a>keyWorld60      = 220;
    <a id="L174"></a>keyWorld61      = 221;
    <a id="L175"></a>keyWorld62      = 222;
    <a id="L176"></a>keyWorld63      = 223;
    <a id="L177"></a>keyWorld64      = 224;
    <a id="L178"></a>keyWorld65      = 225;
    <a id="L179"></a>keyWorld66      = 226;
    <a id="L180"></a>keyWorld67      = 227;
    <a id="L181"></a>keyWorld68      = 228;
    <a id="L182"></a>keyWorld69      = 229;
    <a id="L183"></a>keyWorld70      = 230;
    <a id="L184"></a>keyWorld71      = 231;
    <a id="L185"></a>keyWorld72      = 232;
    <a id="L186"></a>keyWorld73      = 233;
    <a id="L187"></a>keyWorld74      = 234;
    <a id="L188"></a>keyWorld75      = 235;
    <a id="L189"></a>keyWorld76      = 236;
    <a id="L190"></a>keyWorld77      = 237;
    <a id="L191"></a>keyWorld78      = 238;
    <a id="L192"></a>keyWorld79      = 239;
    <a id="L193"></a>keyWorld80      = 240;
    <a id="L194"></a>keyWorld81      = 241;
    <a id="L195"></a>keyWorld82      = 242;
    <a id="L196"></a>keyWorld83      = 243;
    <a id="L197"></a>keyWorld84      = 244;
    <a id="L198"></a>keyWorld85      = 245;
    <a id="L199"></a>keyWorld86      = 246;
    <a id="L200"></a>keyWorld87      = 247;
    <a id="L201"></a>keyWorld88      = 248;
    <a id="L202"></a>keyWorld89      = 249;
    <a id="L203"></a>keyWorld90      = 250;
    <a id="L204"></a>keyWorld91      = 251;
    <a id="L205"></a>keyWorld92      = 252;
    <a id="L206"></a>keyWorld93      = 253;
    <a id="L207"></a>keyWorld94      = 254;
    <a id="L208"></a>keyWorld95      = 255;

    <a id="L210"></a><span class="comment">// Numeric keypad</span>
    <a id="L211"></a>keyKp0        = 256;
    <a id="L212"></a>keyKp1        = 257;
    <a id="L213"></a>keyKp2        = 258;
    <a id="L214"></a>keyKp3        = 259;
    <a id="L215"></a>keyKp4        = 260;
    <a id="L216"></a>keyKp5        = 261;
    <a id="L217"></a>keyKp6        = 262;
    <a id="L218"></a>keyKp7        = 263;
    <a id="L219"></a>keyKp8        = 264;
    <a id="L220"></a>keyKp9        = 265;
    <a id="L221"></a>keyKpPeriod   = 266;
    <a id="L222"></a>keyKpDivide   = 267;
    <a id="L223"></a>keyKpMultiply = 268;
    <a id="L224"></a>keyKpMinus    = 269;
    <a id="L225"></a>keyKpPlus     = 270;
    <a id="L226"></a>keyKpEnter    = 271;
    <a id="L227"></a>keyKpEquals   = 272;

    <a id="L229"></a><span class="comment">// Arrow &amp; insert/delete pad</span>
    <a id="L230"></a>keyUp       = 273;
    <a id="L231"></a>keyDown     = 274;
    <a id="L232"></a>keyRight    = 275;
    <a id="L233"></a>keyLeft     = 276;
    <a id="L234"></a>keyInsert   = 277;
    <a id="L235"></a>keyHome     = 278;
    <a id="L236"></a>keyEnd      = 279;
    <a id="L237"></a>keyPageup   = 280;
    <a id="L238"></a>keyPagedown = 281;

    <a id="L240"></a><span class="comment">// Function keys</span>
    <a id="L241"></a>keyF1  = 282;
    <a id="L242"></a>keyF2  = 283;
    <a id="L243"></a>keyF3  = 284;
    <a id="L244"></a>keyF4  = 285;
    <a id="L245"></a>keyF5  = 286;
    <a id="L246"></a>keyF6  = 287;
    <a id="L247"></a>keyF7  = 288;
    <a id="L248"></a>keyF8  = 289;
    <a id="L249"></a>keyF9  = 290;
    <a id="L250"></a>keyF10 = 291;
    <a id="L251"></a>keyF11 = 292;
    <a id="L252"></a>keyF12 = 293;
    <a id="L253"></a>keyF13 = 294;
    <a id="L254"></a>keyF14 = 295;
    <a id="L255"></a>keyF15 = 296;

    <a id="L257"></a><span class="comment">// Modifier keys</span>
    <a id="L258"></a>keyNumlock   = 300;
    <a id="L259"></a>keyCapslock  = 301;
    <a id="L260"></a>keyScrollock = 302;
    <a id="L261"></a>keyRshift    = 303;
    <a id="L262"></a>keyLshift    = 304;
    <a id="L263"></a>keyRctrl     = 305;
    <a id="L264"></a>keyLctrl     = 306;
    <a id="L265"></a>keyRalt      = 307;
    <a id="L266"></a>keyLalt      = 308;
    <a id="L267"></a>keyRmeta     = 309;
    <a id="L268"></a>keyLmeta     = 310;
    <a id="L269"></a>keyLsuper    = 311;
    <a id="L270"></a>keyRsuper    = 312;
    <a id="L271"></a>keyMode      = 313;
    <a id="L272"></a>keyCompose   = 314;

    <a id="L274"></a><span class="comment">// Misc keys</span>
    <a id="L275"></a>keyHelp   = 315;
    <a id="L276"></a>keyPrint  = 316;
    <a id="L277"></a>keySysreq = 317;
    <a id="L278"></a>keyBreak  = 318;
    <a id="L279"></a>keyMenu   = 319;
    <a id="L280"></a>keyPower  = 320;
    <a id="L281"></a>keyEuro   = 321;
    <a id="L282"></a>keyUndo   = 322;

    <a id="L284"></a><span class="comment">// Add any other keys here</span>
    <a id="L285"></a>keyLast;
<a id="L286"></a>)

<a id="L288"></a><span class="comment">// A keymod is a set of bit flags</span>
<a id="L289"></a>type keymod uint16

<a id="L291"></a>const (
    <a id="L292"></a>keymodNone     = 0x0000;
    <a id="L293"></a>keymodLshift   = 0x0001;
    <a id="L294"></a>keymodRshift   = 0x0002;
    <a id="L295"></a>keymodLctrl    = 0x0040;
    <a id="L296"></a>keymodRctrl    = 0x0080;
    <a id="L297"></a>keymodLalt     = 0x0100;
    <a id="L298"></a>keymodRalt     = 0x0200;
    <a id="L299"></a>keymodLmeta    = 0x0400;
    <a id="L300"></a>keymodRmeta    = 0x0800;
    <a id="L301"></a>keymodNum      = 0x1000;
    <a id="L302"></a>keymodCaps     = 0x2000;
    <a id="L303"></a>keymodMode     = 0x4000;
    <a id="L304"></a>keymodReserved = 0x8000;
<a id="L305"></a>)

<a id="L307"></a>const (
    <a id="L308"></a>mouseButtonLeft   = 1;
    <a id="L309"></a>mouseButtonMiddle = 2;
    <a id="L310"></a>mouseButtonRight  = 3;
    <a id="L311"></a>mouseScrollUp     = 4;
    <a id="L312"></a>mouseScrollDown   = 5;
<a id="L313"></a>)

<a id="L315"></a>const (
    <a id="L316"></a>mouseStateLeftButtonPressed   = 1;
    <a id="L317"></a>mouseStateMiddleButtonPressed = 2;
    <a id="L318"></a>mouseStateRightButtonPressed  = 4;
<a id="L319"></a>)

<a id="L321"></a>const (
    <a id="L322"></a>activeMouse       = 1; <span class="comment">//  mouse leaving/entering</span>
    <a id="L323"></a>activeInputFocus  = 2; <span class="comment">// input focus lost/restored</span>
    <a id="L324"></a>activeApplication = 4; <span class="comment">// application minimized/restored</span>
<a id="L325"></a>)

<a id="L327"></a>const maxEventBytes = 64

<a id="L329"></a>type activeEvent struct {
    <a id="L330"></a>EventType eventType;
    <a id="L331"></a>Gain      uint8;
    <a id="L332"></a>State     uint8;
<a id="L333"></a>}

<a id="L335"></a>type exposeEvent struct {
    <a id="L336"></a>EventType eventType;
<a id="L337"></a>}

<a id="L339"></a>type keyboardEvent struct {
    <a id="L340"></a>EventType eventType;
    <a id="L341"></a>Device    uint8;
    <a id="L342"></a>State     uint8;
    <a id="L343"></a>Pad       uint8;
    <a id="L344"></a>ScanCode  uint8;
    <a id="L345"></a>Pad1      uint8;
    <a id="L346"></a>Key       key;
    <a id="L347"></a>Mod       keymod;
    <a id="L348"></a>Unicode   uint16;
<a id="L349"></a>}

<a id="L351"></a>type mouseMotionEvent struct {
    <a id="L352"></a>EventType eventType;
    <a id="L353"></a>Device    uint8;
    <a id="L354"></a>Buttons   uint8;
    <a id="L355"></a>Pad       uint8;
    <a id="L356"></a>X         uint16;
    <a id="L357"></a>Y         uint16;
    <a id="L358"></a>Xrel      int16;
    <a id="L359"></a>Yrel      int16;
<a id="L360"></a>}

<a id="L362"></a>type mouseButtonEvent struct {
    <a id="L363"></a>EventType eventType;
    <a id="L364"></a>Device    uint8;
    <a id="L365"></a>Button    uint8;
    <a id="L366"></a>State     uint8;
    <a id="L367"></a>X         uint16;
    <a id="L368"></a>Y         uint16;
<a id="L369"></a>}

<a id="L371"></a>type quitEvent struct {
    <a id="L372"></a>EventType eventType;
<a id="L373"></a>}

<a id="L375"></a>type syncEvent struct{}

<a id="L377"></a>type event interface{}

<a id="L379"></a>type reader []byte

<a id="L381"></a>func (r *reader) Read(p []byte) (n int, err os.Error) {
    <a id="L382"></a>b := *r;
    <a id="L383"></a>if len(b) == 0 &amp;&amp; len(p) &gt; 0 {
        <a id="L384"></a>return 0, os.EOF
    <a id="L385"></a>}
    <a id="L386"></a>n = bytes.Copy(p, b);
    <a id="L387"></a>*r = b[n:len(b)];
    <a id="L388"></a>return;
<a id="L389"></a>}

<a id="L391"></a>func (w *Window) readEvents() {
    <a id="L392"></a>buf := make([]byte, maxEventBytes);
    <a id="L393"></a>clean := false;
    <a id="L394"></a>var (
        <a id="L395"></a>ea  *activeEvent;
        <a id="L396"></a>ee  *exposeEvent;
        <a id="L397"></a>ke  *keyboardEvent;
        <a id="L398"></a>mme *mouseMotionEvent;
        <a id="L399"></a>mbe *mouseButtonEvent;
        <a id="L400"></a>qe  *quitEvent;
    <a id="L401"></a>)
    <a id="L402"></a>var m draw.Mouse;
    <a id="L403"></a>for {
        <a id="L404"></a>if err := videoPollEvent(buf); err != nil {
            <a id="L405"></a>if !clean {
                <a id="L406"></a>clean = w.resizec &lt;- false
            <a id="L407"></a>}
            <a id="L408"></a>time.Sleep(10e6); <span class="comment">// 10ms</span>
            <a id="L409"></a>continue;
        <a id="L410"></a>}
        <a id="L411"></a>clean = false;
        <a id="L412"></a>var e event;
        <a id="L413"></a>switch buf[0] {
        <a id="L414"></a>default:
            <a id="L415"></a>log.Stdout(&#34;unsupported event type&#34;, buf[0]);
            <a id="L416"></a>continue;
        <a id="L417"></a>case eventActive:
            <a id="L418"></a>ea = new(activeEvent);
            <a id="L419"></a>e = ea;
        <a id="L420"></a>case eventExpose:
            <a id="L421"></a>ee = new(exposeEvent);
            <a id="L422"></a>e = ee;
        <a id="L423"></a>case eventKeyDown, eventKeyUp:
            <a id="L424"></a>ke = new(keyboardEvent);
            <a id="L425"></a>e = ke;
        <a id="L426"></a>case eventMouseMotion:
            <a id="L427"></a>mme = new(mouseMotionEvent);
            <a id="L428"></a>e = mme;
        <a id="L429"></a>case eventMouseButtonDown, eventMouseButtonUp:
            <a id="L430"></a>mbe = new(mouseButtonEvent);
            <a id="L431"></a>e = mbe;
        <a id="L432"></a>case eventQuit:
            <a id="L433"></a>qe = new(quitEvent);
            <a id="L434"></a>e = qe;
        <a id="L435"></a>}
        <a id="L436"></a>r := reader(buf);
        <a id="L437"></a>if err := binary.Read(&amp;r, binary.LittleEndian, e); err != nil {
            <a id="L438"></a>log.Stdout(&#34;unpacking %T event: %s&#34;, e, err);
            <a id="L439"></a>continue;
        <a id="L440"></a>}
        <a id="L441"></a><span class="comment">// log.Stdoutf(&#34;%#v\n&#34;, e);</span>
        <a id="L442"></a>switch buf[0] {
        <a id="L443"></a>case eventExpose:
            <a id="L444"></a>w.resizec &lt;- true
        <a id="L445"></a>case eventKeyDown:
            <a id="L446"></a>w.kbdc &lt;- int(ke.Key)
        <a id="L447"></a>case eventKeyUp:
            <a id="L448"></a>w.kbdc &lt;- -int(ke.Key)
        <a id="L449"></a>case eventMouseMotion:
            <a id="L450"></a>m.X = int(mme.X);
            <a id="L451"></a>m.Y = int(mme.Y);
            <a id="L452"></a>m.Buttons = int(mme.Buttons);
            <a id="L453"></a>m.Nsec = time.Nanoseconds();
            <a id="L454"></a>_ = w.mousec &lt;- m;
        <a id="L455"></a>case eventMouseButtonDown:
            <a id="L456"></a>m.X = int(mbe.X);
            <a id="L457"></a>m.Y = int(mbe.Y);
            <a id="L458"></a><span class="comment">// TODO(rsc): Remove uint cast once 8g bug is fixed.</span>
            <a id="L459"></a>m.Buttons |= 1 &lt;&lt; uint(mbe.Button-1);
            <a id="L460"></a>m.Nsec = time.Nanoseconds();
            <a id="L461"></a>_ = w.mousec &lt;- m;
        <a id="L462"></a>case eventMouseButtonUp:
            <a id="L463"></a>m.X = int(mbe.X);
            <a id="L464"></a>m.Y = int(mbe.Y);
            <a id="L465"></a><span class="comment">// TODO(rsc): Remove uint cast once 8g bug is fixed.</span>
            <a id="L466"></a>m.Buttons &amp;^= 1 &lt;&lt; uint(mbe.Button-1);
            <a id="L467"></a>m.Nsec = time.Nanoseconds();
            <a id="L468"></a>_ = w.mousec &lt;- m;
        <a id="L469"></a>case eventQuit:
            <a id="L470"></a>w.quitc &lt;- true
        <a id="L471"></a>}
    <a id="L472"></a>}
<a id="L473"></a>}
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
