<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/4s/xs.go</title>

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
	<li>Sun Nov 15 20:28:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/4s/xs.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// games/4s - a tetris clone</span>
<a id="L2"></a><span class="comment">//</span>
<a id="L3"></a><span class="comment">// Derived from Plan 9&#39;s /sys/src/games/xs.c</span>
<a id="L4"></a><span class="comment">// http://plan9.bell-labs.com/sources/plan9/sys/src/games/xs.c</span>
<a id="L5"></a><span class="comment">//</span>
<a id="L6"></a><span class="comment">// Copyright (C) 2003, Lucent Technologies Inc. and others. All Rights Reserved.</span>
<a id="L7"></a><span class="comment">// Portions Copyright 2009 The Go Authors.  All Rights Reserved.</span>
<a id="L8"></a><span class="comment">// Distributed under the terms of the Lucent Public License Version 1.02</span>
<a id="L9"></a><span class="comment">// See http://plan9.bell-labs.com/plan9/license.html</span>

<a id="L11"></a><span class="comment">/*</span>
<a id="L12"></a><span class="comment"> * engine for 4s, 5s, etc</span>
<a id="L13"></a><span class="comment"> */</span>

<a id="L15"></a>package main

<a id="L17"></a>import (
    <a id="L18"></a>&#34;exp/draw&#34;;
    <a id="L19"></a>&#34;image&#34;;
    <a id="L20"></a>&#34;log&#34;;
    <a id="L21"></a>&#34;os&#34;;
    <a id="L22"></a>&#34;rand&#34;;
    <a id="L23"></a>&#34;time&#34;;
<a id="L24"></a>)

<a id="L26"></a><span class="comment">/*</span>
<a id="L27"></a><span class="comment">Cursor whitearrow = {</span>
<a id="L28"></a><span class="comment">	{0, 0},</span>
<a id="L29"></a><span class="comment">	{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE, 0xFF, 0xFC,</span>
<a id="L30"></a><span class="comment">	 0xFF, 0xF0, 0xFF, 0xF0, 0xFF, 0xF8, 0xFF, 0xFC,</span>
<a id="L31"></a><span class="comment">	 0xFF, 0xFE, 0xFF, 0xFF, 0xFF, 0xFE, 0xFF, 0xFC,</span>
<a id="L32"></a><span class="comment">	 0xF3, 0xF8, 0xF1, 0xF0, 0xE0, 0xE0, 0xC0, 0x40, },</span>
<a id="L33"></a><span class="comment">	{0xFF, 0xFF, 0xFF, 0xFF, 0xC0, 0x06, 0xC0, 0x1C,</span>
<a id="L34"></a><span class="comment">	 0xC0, 0x30, 0xC0, 0x30, 0xC0, 0x38, 0xC0, 0x1C,</span>
<a id="L35"></a><span class="comment">	 0xC0, 0x0E, 0xC0, 0x07, 0xCE, 0x0E, 0xDF, 0x1C,</span>
<a id="L36"></a><span class="comment">	 0xD3, 0xB8, 0xF1, 0xF0, 0xE0, 0xE0, 0xC0, 0x40, }</span>
<a id="L37"></a><span class="comment">};</span>
<a id="L38"></a><span class="comment">*/</span>

<a id="L40"></a>const (
    <a id="L41"></a>CNone   = 0;
    <a id="L42"></a>CBounds = 1;
    <a id="L43"></a>CPiece  = 2;
    <a id="L44"></a>NX      = 10;
    <a id="L45"></a>NY      = 20;

    <a id="L47"></a>NCOL = 10;

    <a id="L49"></a>MAXN = 5;
<a id="L50"></a>)

<a id="L52"></a>var (
    <a id="L53"></a>N                        int;
    <a id="L54"></a>display                  draw.Context;
    <a id="L55"></a>screen                   draw.Image;
    <a id="L56"></a>screenr                  draw.Rectangle;
    <a id="L57"></a>board                    [NY][NX]byte;
    <a id="L58"></a>rboard                   draw.Rectangle;
    <a id="L59"></a>pscore                   draw.Point;
    <a id="L60"></a>scoresz                  draw.Point;
    <a id="L61"></a>pcsz                     = 32;
    <a id="L62"></a>pos                      draw.Point;
    <a id="L63"></a>bbr, bb2r                draw.Rectangle;
    <a id="L64"></a>bb, bbmask, bb2, bb2mask *image.RGBA;
    <a id="L65"></a>whitemask                image.Image;
    <a id="L66"></a>br, br2                  draw.Rectangle;
    <a id="L67"></a>points                   int;
    <a id="L68"></a>dt                       int;
    <a id="L69"></a>DY                       int;
    <a id="L70"></a>DMOUSE                   int;
    <a id="L71"></a>lastmx                   int;
    <a id="L72"></a>mouse                    draw.Mouse;
    <a id="L73"></a>newscreen                bool;
    <a id="L74"></a>timerc                   &lt;-chan int64;
    <a id="L75"></a>suspc                    chan bool;
    <a id="L76"></a>mousec                   chan draw.Mouse;
    <a id="L77"></a>resizec                  &lt;-chan bool;
    <a id="L78"></a>kbdc                     chan int;
    <a id="L79"></a>suspended                bool;
    <a id="L80"></a>tsleep                   int;
    <a id="L81"></a>piece                    *Piece;
    <a id="L82"></a>pieces                   []Piece;
<a id="L83"></a>)

<a id="L85"></a>type Piece struct {
    <a id="L86"></a>rot   int;
    <a id="L87"></a>tx    int;
    <a id="L88"></a>sz    draw.Point;
    <a id="L89"></a>d     []draw.Point;
    <a id="L90"></a>left  *Piece;
    <a id="L91"></a>right *Piece;
<a id="L92"></a>}

<a id="L94"></a>var txbits = [NCOL][32]byte{
    <a id="L95"></a>[32]byte{0xDD, 0xDD, 0xFF, 0xFF, 0x77, 0x77, 0xFF, 0xFF,
        <a id="L96"></a>0xDD, 0xDD, 0xFF, 0xFF, 0x77, 0x77, 0xFF, 0xFF,
        <a id="L97"></a>0xDD, 0xDD, 0xFF, 0xFF, 0x77, 0x77, 0xFF, 0xFF,
        <a id="L98"></a>0xDD, 0xDD, 0xFF, 0xFF, 0x77, 0x77, 0xFF, 0xFF,
    <a id="L99"></a>},
    <a id="L100"></a>[32]byte{0xDD, 0xDD, 0x77, 0x77, 0xDD, 0xDD, 0x77, 0x77,
        <a id="L101"></a>0xDD, 0xDD, 0x77, 0x77, 0xDD, 0xDD, 0x77, 0x77,
        <a id="L102"></a>0xDD, 0xDD, 0x77, 0x77, 0xDD, 0xDD, 0x77, 0x77,
        <a id="L103"></a>0xDD, 0xDD, 0x77, 0x77, 0xDD, 0xDD, 0x77, 0x77,
    <a id="L104"></a>},
    <a id="L105"></a>[32]byte{0xAA, 0xAA, 0x55, 0x55, 0xAA, 0xAA, 0x55, 0x55,
        <a id="L106"></a>0xAA, 0xAA, 0x55, 0x55, 0xAA, 0xAA, 0x55, 0x55,
        <a id="L107"></a>0xAA, 0xAA, 0x55, 0x55, 0xAA, 0xAA, 0x55, 0x55,
        <a id="L108"></a>0xAA, 0xAA, 0x55, 0x55, 0xAA, 0xAA, 0x55, 0x55,
    <a id="L109"></a>},
    <a id="L110"></a>[32]byte{0xAA, 0xAA, 0x55, 0x55, 0xAA, 0xAA, 0x55, 0x55,
        <a id="L111"></a>0xAA, 0xAA, 0x55, 0x55, 0xAA, 0xAA, 0x55, 0x55,
        <a id="L112"></a>0xAA, 0xAA, 0x55, 0x55, 0xAA, 0xAA, 0x55, 0x55,
        <a id="L113"></a>0xAA, 0xAA, 0x55, 0x55, 0xAA, 0xAA, 0x55, 0x55,
    <a id="L114"></a>},
    <a id="L115"></a>[32]byte{0x22, 0x22, 0x88, 0x88, 0x22, 0x22, 0x88, 0x88,
        <a id="L116"></a>0x22, 0x22, 0x88, 0x88, 0x22, 0x22, 0x88, 0x88,
        <a id="L117"></a>0x22, 0x22, 0x88, 0x88, 0x22, 0x22, 0x88, 0x88,
        <a id="L118"></a>0x22, 0x22, 0x88, 0x88, 0x22, 0x22, 0x88, 0x88,
    <a id="L119"></a>},
    <a id="L120"></a>[32]byte{0x22, 0x22, 0x00, 0x00, 0x88, 0x88, 0x00, 0x00,
        <a id="L121"></a>0x22, 0x22, 0x00, 0x00, 0x88, 0x88, 0x00, 0x00,
        <a id="L122"></a>0x22, 0x22, 0x00, 0x00, 0x88, 0x88, 0x00, 0x00,
        <a id="L123"></a>0x22, 0x22, 0x00, 0x00, 0x88, 0x88, 0x00, 0x00,
    <a id="L124"></a>},
    <a id="L125"></a>[32]byte{0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00,
        <a id="L126"></a>0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00,
        <a id="L127"></a>0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00,
        <a id="L128"></a>0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00,
    <a id="L129"></a>},
    <a id="L130"></a>[32]byte{0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00,
        <a id="L131"></a>0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00,
        <a id="L132"></a>0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00,
        <a id="L133"></a>0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00,
    <a id="L134"></a>},
    <a id="L135"></a>[32]byte{0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC,
        <a id="L136"></a>0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC,
        <a id="L137"></a>0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC,
        <a id="L138"></a>0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC,
    <a id="L139"></a>},
    <a id="L140"></a>[32]byte{0xCC, 0xCC, 0xCC, 0xCC, 0x33, 0x33, 0x33, 0x33,
        <a id="L141"></a>0xCC, 0xCC, 0xCC, 0xCC, 0x33, 0x33, 0x33, 0x33,
        <a id="L142"></a>0xCC, 0xCC, 0xCC, 0xCC, 0x33, 0x33, 0x33, 0x33,
        <a id="L143"></a>0xCC, 0xCC, 0xCC, 0xCC, 0x33, 0x33, 0x33, 0x33,
    <a id="L144"></a>},
<a id="L145"></a>}

<a id="L147"></a>var txpix = [NCOL]draw.Color{
    <a id="L148"></a>draw.Yellow,            <span class="comment">/* yellow */</span>
    <a id="L149"></a>draw.Cyan,              <span class="comment">/* cyan */</span>
    <a id="L150"></a>draw.Green,             <span class="comment">/* lime green */</span>
    <a id="L151"></a>draw.GreyBlue,          <span class="comment">/* slate */</span>
    <a id="L152"></a>draw.Red,               <span class="comment">/* red */</span>
    <a id="L153"></a>draw.GreyGreen,         <span class="comment">/* olive green */</span>
    <a id="L154"></a>draw.Blue,              <span class="comment">/* blue */</span>
    <a id="L155"></a>draw.Color(0xFF55AAFF), <span class="comment">/* pink */</span>
    <a id="L156"></a>draw.Color(0xFFAAFFFF), <span class="comment">/* lavender */</span>
    <a id="L157"></a>draw.Color(0xBB005DFF), <span class="comment">/* maroon */</span>
<a id="L158"></a>}

<a id="L160"></a>func movemouse() int {
    <a id="L161"></a><span class="comment">//mouse.draw.Point = draw.Pt(rboard.Min.X + rboard.Dx()/2, rboard.Min.Y + rboard.Dy()/2);</span>
    <a id="L162"></a><span class="comment">//moveto(mousectl, mouse.Xy);</span>
    <a id="L163"></a>return mouse.X
<a id="L164"></a>}

<a id="L166"></a>func warp(p draw.Point, x int) int {
    <a id="L167"></a>if !suspended &amp;&amp; piece != nil {
        <a id="L168"></a>x = pos.X + piece.sz.X*pcsz/2;
        <a id="L169"></a>if p.Y &lt; rboard.Min.Y {
            <a id="L170"></a>p.Y = rboard.Min.Y
        <a id="L171"></a>}
        <a id="L172"></a>if p.Y &gt;= rboard.Max.Y {
            <a id="L173"></a>p.Y = rboard.Max.Y - 1
        <a id="L174"></a>}
        <a id="L175"></a><span class="comment">//moveto(mousectl, draw.Pt(x, p.Y));</span>
    <a id="L176"></a>}
    <a id="L177"></a>return x;
<a id="L178"></a>}

<a id="L180"></a>func initPieces() {
    <a id="L181"></a>for i := range pieces {
        <a id="L182"></a>p := &amp;pieces[i];
        <a id="L183"></a>if p.rot == 3 {
            <a id="L184"></a>p.right = &amp;pieces[i-3]
        <a id="L185"></a>} else {
            <a id="L186"></a>p.right = &amp;pieces[i+1]
        <a id="L187"></a>}
        <a id="L188"></a>if p.rot == 0 {
            <a id="L189"></a>p.left = &amp;pieces[i+3]
        <a id="L190"></a>} else {
            <a id="L191"></a>p.left = &amp;pieces[i-1]
        <a id="L192"></a>}
    <a id="L193"></a>}
<a id="L194"></a>}

<a id="L196"></a>func collide(pt draw.Point, p *Piece) bool {
    <a id="L197"></a>pt.X = (pt.X - rboard.Min.X) / pcsz;
    <a id="L198"></a>pt.Y = (pt.Y - rboard.Min.Y) / pcsz;
    <a id="L199"></a>for _, q := range p.d {
        <a id="L200"></a>pt.X += q.X;
        <a id="L201"></a>pt.Y += q.Y;
        <a id="L202"></a>if pt.X &lt; 0 || pt.X &gt;= NX || pt.Y &lt; 0 || pt.Y &gt;= NY {
            <a id="L203"></a>return true;
            <a id="L204"></a>continue;
        <a id="L205"></a>}
        <a id="L206"></a>if board[pt.Y][pt.X] != 0 {
            <a id="L207"></a>return true
        <a id="L208"></a>}
    <a id="L209"></a>}
    <a id="L210"></a>return false;
<a id="L211"></a>}

<a id="L213"></a>func collider(pt, pmax draw.Point) bool {
    <a id="L214"></a>pi := (pt.X - rboard.Min.X) / pcsz;
    <a id="L215"></a>pj := (pt.Y - rboard.Min.Y) / pcsz;
    <a id="L216"></a>n := pmax.X / pcsz;
    <a id="L217"></a>m := pmax.Y/pcsz + 1;
    <a id="L218"></a>for i := pi; i &lt; pi+n &amp;&amp; i &lt; NX; i++ {
        <a id="L219"></a>for j := pj; j &lt; pj+m &amp;&amp; j &lt; NY; j++ {
            <a id="L220"></a>if board[j][i] != 0 {
                <a id="L221"></a>return true
            <a id="L222"></a>}
        <a id="L223"></a>}
    <a id="L224"></a>}
    <a id="L225"></a>return false;
<a id="L226"></a>}

<a id="L228"></a>func setpiece(p *Piece) {
    <a id="L229"></a>draw.Draw(bb, bbr, draw.White, nil, draw.ZP);
    <a id="L230"></a>draw.Draw(bbmask, bbr, draw.Transparent, nil, draw.ZP);
    <a id="L231"></a>br = draw.Rect(0, 0, 0, 0);
    <a id="L232"></a>br2 = br;
    <a id="L233"></a>piece = p;
    <a id="L234"></a>if p == nil {
        <a id="L235"></a>return
    <a id="L236"></a>}
    <a id="L237"></a>var op draw.Point;
    <a id="L238"></a>var r draw.Rectangle;
    <a id="L239"></a>r.Min = bbr.Min;
    <a id="L240"></a>for i, pt := range p.d {
        <a id="L241"></a>r.Min.X += pt.X * pcsz;
        <a id="L242"></a>r.Min.Y += pt.Y * pcsz;
        <a id="L243"></a>r.Max.X = r.Min.X + pcsz;
        <a id="L244"></a>r.Max.Y = r.Min.Y + pcsz;
        <a id="L245"></a>if i == 0 {
            <a id="L246"></a>draw.Draw(bb, r, draw.Black, nil, draw.ZP);
            <a id="L247"></a>draw.Draw(bb, r.Inset(1), txpix[piece.tx], nil, draw.ZP);
            <a id="L248"></a>draw.Draw(bbmask, r, draw.Opaque, nil, draw.ZP);
            <a id="L249"></a>op = r.Min;
        <a id="L250"></a>} else {
            <a id="L251"></a>draw.Draw(bb, r, bb, nil, op);
            <a id="L252"></a>draw.Draw(bbmask, r, bbmask, nil, op);
        <a id="L253"></a>}
        <a id="L254"></a>if br.Max.X &lt; r.Max.X {
            <a id="L255"></a>br.Max.X = r.Max.X
        <a id="L256"></a>}
        <a id="L257"></a>if br.Max.Y &lt; r.Max.Y {
            <a id="L258"></a>br.Max.Y = r.Max.Y
        <a id="L259"></a>}
    <a id="L260"></a>}
    <a id="L261"></a>br.Max = br.Max.Sub(bbr.Min);
    <a id="L262"></a>delta := draw.Pt(0, DY);
    <a id="L263"></a>br2.Max = br.Max.Add(delta);
    <a id="L264"></a>r = br.Add(bb2r.Min);
    <a id="L265"></a>r2 := br2.Add(bb2r.Min);
    <a id="L266"></a>draw.Draw(bb2, r2, draw.White, nil, draw.ZP);
    <a id="L267"></a>draw.Draw(bb2, r.Add(delta), bb, nil, bbr.Min);
    <a id="L268"></a>draw.Draw(bb2mask, r2, draw.Transparent, nil, draw.ZP);
    <a id="L269"></a>draw.Draw(bb2mask, r, draw.Opaque, bbmask, bbr.Min);
    <a id="L270"></a>draw.Draw(bb2mask, r.Add(delta), draw.Opaque, bbmask, bbr.Min);
<a id="L271"></a>}

<a id="L273"></a>func drawpiece() {
    <a id="L274"></a>draw.Draw(screen, br.Add(pos), bb, bbmask, bbr.Min);
    <a id="L275"></a>if suspended {
        <a id="L276"></a>draw.Draw(screen, br.Add(pos), draw.White, whitemask, draw.ZP)
    <a id="L277"></a>}
<a id="L278"></a>}

<a id="L280"></a>func undrawpiece() {
    <a id="L281"></a>var mask image.Image;
    <a id="L282"></a>if collider(pos, br.Max) {
        <a id="L283"></a>mask = bbmask
    <a id="L284"></a>}
    <a id="L285"></a>draw.Draw(screen, br.Add(pos), draw.White, mask, bbr.Min);
<a id="L286"></a>}

<a id="L288"></a>func rest() {
    <a id="L289"></a>pt := pos.Sub(rboard.Min).Div(pcsz);
    <a id="L290"></a>for _, p := range piece.d {
        <a id="L291"></a>pt.X += p.X;
        <a id="L292"></a>pt.Y += p.Y;
        <a id="L293"></a>board[pt.Y][pt.X] = byte(piece.tx + 16);
    <a id="L294"></a>}
<a id="L295"></a>}

<a id="L297"></a>func canfit(p *Piece) bool {
    <a id="L298"></a>var dx = [...]int{0, -1, 1, -2, 2, -3, 3, 4, -4};
    <a id="L299"></a>j := N + 1;
    <a id="L300"></a>if j &gt;= 4 {
        <a id="L301"></a>j = p.sz.X;
        <a id="L302"></a>if j &lt; p.sz.Y {
            <a id="L303"></a>j = p.sz.Y
        <a id="L304"></a>}
        <a id="L305"></a>j = 2*j - 1;
    <a id="L306"></a>}
    <a id="L307"></a>for i := 0; i &lt; j; i++ {
        <a id="L308"></a>var z draw.Point;
        <a id="L309"></a>z.X = pos.X + dx[i]*pcsz;
        <a id="L310"></a>z.Y = pos.Y;
        <a id="L311"></a>if !collide(z, p) {
            <a id="L312"></a>z.Y = pos.Y + pcsz - 1;
            <a id="L313"></a>if !collide(z, p) {
                <a id="L314"></a>undrawpiece();
                <a id="L315"></a>pos.X = z.X;
                <a id="L316"></a>return true;
            <a id="L317"></a>}
        <a id="L318"></a>}
    <a id="L319"></a>}
    <a id="L320"></a>return false;
<a id="L321"></a>}

<a id="L323"></a>func score(p int) {
    <a id="L324"></a>points += p
    <a id="L325"></a><span class="comment">//	snprint(buf, sizeof(buf), &#34;%.6ld&#34;, points);</span>
    <a id="L326"></a><span class="comment">//	draw.Draw(screen, draw.Rpt(pscore, pscore.Add(scoresz)), draw.White, nil, draw.ZP);</span>
    <a id="L327"></a><span class="comment">//	string(screen, pscore, draw.Black, draw.ZP, font, buf);</span>
<a id="L328"></a>}

<a id="L330"></a>func drawsq(b draw.Image, p draw.Point, ptx int) {
    <a id="L331"></a>var r draw.Rectangle;
    <a id="L332"></a>r.Min = p;
    <a id="L333"></a>r.Max.X = r.Min.X + pcsz;
    <a id="L334"></a>r.Max.Y = r.Min.Y + pcsz;
    <a id="L335"></a>draw.Draw(b, r, draw.Black, nil, draw.ZP);
    <a id="L336"></a>draw.Draw(b, r.Inset(1), txpix[ptx], nil, draw.ZP);
<a id="L337"></a>}

<a id="L339"></a>func drawboard() {
    <a id="L340"></a>draw.Border(screen, rboard.Inset(-2), 2, draw.Black, draw.ZP);
    <a id="L341"></a>draw.Draw(screen, draw.Rect(rboard.Min.X, rboard.Min.Y-2, rboard.Max.X, rboard.Min.Y),
        <a id="L342"></a>draw.White, nil, draw.ZP);
    <a id="L343"></a>for i := 0; i &lt; NY; i++ {
        <a id="L344"></a>for j := 0; j &lt; NX; j++ {
            <a id="L345"></a>if board[i][j] != 0 {
                <a id="L346"></a>drawsq(screen, draw.Pt(rboard.Min.X+j*pcsz, rboard.Min.Y+i*pcsz), int(board[i][j]-16))
            <a id="L347"></a>}
        <a id="L348"></a>}
    <a id="L349"></a>}
    <a id="L350"></a>score(0);
    <a id="L351"></a>if suspended {
        <a id="L352"></a>draw.Draw(screen, screenr, draw.White, whitemask, draw.ZP)
    <a id="L353"></a>}
<a id="L354"></a>}

<a id="L356"></a>func choosepiece() {
    <a id="L357"></a>for {
        <a id="L358"></a>i := rand.Intn(len(pieces));
        <a id="L359"></a>setpiece(&amp;pieces[i]);
        <a id="L360"></a>pos = rboard.Min;
        <a id="L361"></a>pos.X += rand.Intn(NX) * pcsz;
        <a id="L362"></a>if !collide(draw.Pt(pos.X, pos.Y+pcsz-DY), piece) {
            <a id="L363"></a>break
        <a id="L364"></a>}
    <a id="L365"></a>}
    <a id="L366"></a>drawpiece();
    <a id="L367"></a>display.FlushImage();
<a id="L368"></a>}

<a id="L370"></a>func movepiece() bool {
    <a id="L371"></a>var mask image.Image;
    <a id="L372"></a>if collide(draw.Pt(pos.X, pos.Y+pcsz), piece) {
        <a id="L373"></a>return false
    <a id="L374"></a>}
    <a id="L375"></a>if collider(pos, br2.Max) {
        <a id="L376"></a>mask = bb2mask
    <a id="L377"></a>}
    <a id="L378"></a>draw.Draw(screen, br2.Add(pos), bb2, mask, bb2r.Min);
    <a id="L379"></a>pos.Y += DY;
    <a id="L380"></a>display.FlushImage();
    <a id="L381"></a>return true;
<a id="L382"></a>}

<a id="L384"></a>func suspend(s bool) {
    <a id="L385"></a>suspended = s;
    <a id="L386"></a><span class="comment">/*</span>
    <a id="L387"></a><span class="comment">	if suspended {</span>
    <a id="L388"></a><span class="comment">		setcursor(mousectl, &amp;whitearrow);</span>
    <a id="L389"></a><span class="comment">	} else {</span>
    <a id="L390"></a><span class="comment">		setcursor(mousectl, nil);</span>
    <a id="L391"></a><span class="comment">	}</span>
    <a id="L392"></a><span class="comment">*/</span>
    <a id="L393"></a>if !suspended {
        <a id="L394"></a>drawpiece()
    <a id="L395"></a>}
    <a id="L396"></a>drawboard();
    <a id="L397"></a>display.FlushImage();
<a id="L398"></a>}

<a id="L400"></a>func pause(t int) {
    <a id="L401"></a>display.FlushImage();
    <a id="L402"></a>for {
        <a id="L403"></a>select {
        <a id="L404"></a>case s := &lt;-suspc:
            <a id="L405"></a>if !suspended &amp;&amp; s {
                <a id="L406"></a>suspend(true)
            <a id="L407"></a>} else if suspended &amp;&amp; !s {
                <a id="L408"></a>suspend(false);
                <a id="L409"></a>lastmx = warp(mouse.Point, lastmx);
            <a id="L410"></a>}
        <a id="L411"></a>case &lt;-timerc:
            <a id="L412"></a>if suspended {
                <a id="L413"></a>break
            <a id="L414"></a>}
            <a id="L415"></a>t -= tsleep;
            <a id="L416"></a>if t &lt; 0 {
                <a id="L417"></a>return
            <a id="L418"></a>}
        <a id="L419"></a>case &lt;-resizec:
            <a id="L420"></a><span class="comment">//redraw(true);</span>
        <a id="L421"></a>case mouse = &lt;-mousec:
        <a id="L422"></a>case &lt;-kbdc:
        <a id="L423"></a>}
    <a id="L424"></a>}
<a id="L425"></a>}

<a id="L427"></a>func horiz() bool {
    <a id="L428"></a>var lev [MAXN]int;
    <a id="L429"></a>h := 0;
    <a id="L430"></a>for i := 0; i &lt; NY; i++ {
        <a id="L431"></a>for j := 0; board[i][j] != 0; j++ {
            <a id="L432"></a>if j == NX-1 {
                <a id="L433"></a>lev[h] = i;
                <a id="L434"></a>h++;
                <a id="L435"></a>break;
            <a id="L436"></a>}
        <a id="L437"></a>}
    <a id="L438"></a>}
    <a id="L439"></a>if h == 0 {
        <a id="L440"></a>return false
    <a id="L441"></a>}
    <a id="L442"></a>r := rboard;
    <a id="L443"></a>newscreen = false;
    <a id="L444"></a>for j := 0; j &lt; h; j++ {
        <a id="L445"></a>r.Min.Y = rboard.Min.Y + lev[j]*pcsz;
        <a id="L446"></a>r.Max.Y = r.Min.Y + pcsz;
        <a id="L447"></a>draw.Draw(screen, r, draw.White, whitemask, draw.ZP);
        <a id="L448"></a>display.FlushImage();
    <a id="L449"></a>}
    <a id="L450"></a>PlaySound(whoosh);
    <a id="L451"></a>for i := 0; i &lt; 3; i++ {
        <a id="L452"></a>pause(250);
        <a id="L453"></a>if newscreen {
            <a id="L454"></a>drawboard();
            <a id="L455"></a>break;
        <a id="L456"></a>}
        <a id="L457"></a>for j := 0; j &lt; h; j++ {
            <a id="L458"></a>r.Min.Y = rboard.Min.Y + lev[j]*pcsz;
            <a id="L459"></a>r.Max.Y = r.Min.Y + pcsz;
            <a id="L460"></a>draw.Draw(screen, r, draw.White, whitemask, draw.ZP);
        <a id="L461"></a>}
        <a id="L462"></a>display.FlushImage();
    <a id="L463"></a>}
    <a id="L464"></a>r = rboard;
    <a id="L465"></a>for j := 0; j &lt; h; j++ {
        <a id="L466"></a>i := NY - lev[j] - 1;
        <a id="L467"></a>score(250 + 10*i*i);
        <a id="L468"></a>r.Min.Y = rboard.Min.Y;
        <a id="L469"></a>r.Max.Y = rboard.Min.Y + lev[j]*pcsz;
        <a id="L470"></a>draw.Draw(screen, r.Add(draw.Pt(0, pcsz)), screen, nil, r.Min);
        <a id="L471"></a>r.Max.Y = rboard.Min.Y + pcsz;
        <a id="L472"></a>draw.Draw(screen, r, draw.White, nil, draw.ZP);
        <a id="L473"></a>for k := lev[j] - 1; k &gt;= 0; k-- {
            <a id="L474"></a>board[k+1] = board[k]
        <a id="L475"></a>}
        <a id="L476"></a>board[0] = [NX]byte{};
    <a id="L477"></a>}
    <a id="L478"></a>display.FlushImage();
    <a id="L479"></a>return true;
<a id="L480"></a>}

<a id="L482"></a>func mright() {
    <a id="L483"></a>if !collide(draw.Pt(pos.X+pcsz, pos.Y), piece) &amp;&amp;
        <a id="L484"></a>!collide(draw.Pt(pos.X+pcsz, pos.Y+pcsz-DY), piece) {
        <a id="L485"></a>undrawpiece();
        <a id="L486"></a>pos.X += pcsz;
        <a id="L487"></a>drawpiece();
        <a id="L488"></a>display.FlushImage();
    <a id="L489"></a>}
<a id="L490"></a>}

<a id="L492"></a>func mleft() {
    <a id="L493"></a>if !collide(draw.Pt(pos.X-pcsz, pos.Y), piece) &amp;&amp;
        <a id="L494"></a>!collide(draw.Pt(pos.X-pcsz, pos.Y+pcsz-DY), piece) {
        <a id="L495"></a>undrawpiece();
        <a id="L496"></a>pos.X -= pcsz;
        <a id="L497"></a>drawpiece();
        <a id="L498"></a>display.FlushImage();
    <a id="L499"></a>}
<a id="L500"></a>}

<a id="L502"></a>func rright() {
    <a id="L503"></a>if canfit(piece.right) {
        <a id="L504"></a>setpiece(piece.right);
        <a id="L505"></a>drawpiece();
        <a id="L506"></a>display.FlushImage();
    <a id="L507"></a>}
<a id="L508"></a>}

<a id="L510"></a>func rleft() {
    <a id="L511"></a>if canfit(piece.left) {
        <a id="L512"></a>setpiece(piece.left);
        <a id="L513"></a>drawpiece();
        <a id="L514"></a>display.FlushImage();
    <a id="L515"></a>}
<a id="L516"></a>}

<a id="L518"></a>var fusst = 0

<a id="L520"></a>func drop(f bool) bool {
    <a id="L521"></a>if f {
        <a id="L522"></a>score(5 * (rboard.Max.Y - pos.Y) / pcsz);
        <a id="L523"></a>for movepiece() {
        <a id="L524"></a>}
    <a id="L525"></a>}
    <a id="L526"></a>fusst = 0;
    <a id="L527"></a>rest();
    <a id="L528"></a>if pos.Y == rboard.Min.Y &amp;&amp; !horiz() {
        <a id="L529"></a>return true
    <a id="L530"></a>}
    <a id="L531"></a>horiz();
    <a id="L532"></a>setpiece(nil);
    <a id="L533"></a>pause(1500);
    <a id="L534"></a>choosepiece();
    <a id="L535"></a>lastmx = warp(mouse.Point, lastmx);
    <a id="L536"></a>return false;
<a id="L537"></a>}

<a id="L539"></a>func play() {
    <a id="L540"></a>var om draw.Mouse;
    <a id="L541"></a>dt = 64;
    <a id="L542"></a>lastmx = -1;
    <a id="L543"></a>lastmx = movemouse();
    <a id="L544"></a>choosepiece();
    <a id="L545"></a>lastmx = warp(mouse.Point, lastmx);
    <a id="L546"></a>for {
        <a id="L547"></a>select {
        <a id="L548"></a>case mouse = &lt;-mousec:
            <a id="L549"></a>if suspended {
                <a id="L550"></a>om = mouse;
                <a id="L551"></a>break;
            <a id="L552"></a>}
            <a id="L553"></a>if lastmx &lt; 0 {
                <a id="L554"></a>lastmx = mouse.X
            <a id="L555"></a>}
            <a id="L556"></a>if mouse.X &gt; lastmx+DMOUSE {
                <a id="L557"></a>mright();
                <a id="L558"></a>lastmx = mouse.X;
            <a id="L559"></a>}
            <a id="L560"></a>if mouse.X &lt; lastmx-DMOUSE {
                <a id="L561"></a>mleft();
                <a id="L562"></a>lastmx = mouse.X;
            <a id="L563"></a>}
            <a id="L564"></a>if mouse.Buttons&amp;^om.Buttons&amp;1 == 1 {
                <a id="L565"></a>rleft()
            <a id="L566"></a>}
            <a id="L567"></a>if mouse.Buttons&amp;^om.Buttons&amp;2 == 2 {
                <a id="L568"></a>if drop(true) {
                    <a id="L569"></a>return
                <a id="L570"></a>}
            <a id="L571"></a>}
            <a id="L572"></a>if mouse.Buttons&amp;^om.Buttons&amp;4 == 4 {
                <a id="L573"></a>rright()
            <a id="L574"></a>}
            <a id="L575"></a>om = mouse;

        <a id="L577"></a>case s := &lt;-suspc:
            <a id="L578"></a>if !suspended &amp;&amp; s {
                <a id="L579"></a>suspend(true)
            <a id="L580"></a>} else if suspended &amp;&amp; !s {
                <a id="L581"></a>suspend(false);
                <a id="L582"></a>lastmx = warp(mouse.Point, lastmx);
            <a id="L583"></a>}

        <a id="L585"></a>case &lt;-resizec:
            <a id="L586"></a><span class="comment">//redraw(true);</span>

        <a id="L588"></a>case r := &lt;-kbdc:
            <a id="L589"></a>if suspended {
                <a id="L590"></a>break
            <a id="L591"></a>}
            <a id="L592"></a>switch r {
            <a id="L593"></a>case &#39;f&#39;, &#39;;&#39;:
                <a id="L594"></a>mright()
            <a id="L595"></a>case &#39;a&#39;, &#39;j&#39;:
                <a id="L596"></a>mleft()
            <a id="L597"></a>case &#39;d&#39;, &#39;l&#39;:
                <a id="L598"></a>rright()
            <a id="L599"></a>case &#39;s&#39;, &#39;k&#39;:
                <a id="L600"></a>rleft()
            <a id="L601"></a>case &#39; &#39;:
                <a id="L602"></a>if drop(true) {
                    <a id="L603"></a>return
                <a id="L604"></a>}
            <a id="L605"></a>}

        <a id="L607"></a>case &lt;-timerc:
            <a id="L608"></a>if suspended {
                <a id="L609"></a>break
            <a id="L610"></a>}
            <a id="L611"></a>dt -= tsleep;
            <a id="L612"></a>if dt &lt; 0 {
                <a id="L613"></a>i := 1;
                <a id="L614"></a>dt = 16 * (points + rand.Intn(10000) - 5000) / 10000;
                <a id="L615"></a>if dt &gt;= 32 {
                    <a id="L616"></a>i += (dt - 32) / 16;
                    <a id="L617"></a>dt = 32;
                <a id="L618"></a>}
                <a id="L619"></a>dt = 52 - dt;
                <a id="L620"></a>for ; i &gt; 0; i-- {
                    <a id="L621"></a>if movepiece() {
                        <a id="L622"></a>continue
                    <a id="L623"></a>}
                    <a id="L624"></a>fusst++;
                    <a id="L625"></a>if fusst == 40 {
                        <a id="L626"></a>if drop(false) {
                            <a id="L627"></a>return
                        <a id="L628"></a>}
                        <a id="L629"></a>break;
                    <a id="L630"></a>}
                <a id="L631"></a>}
            <a id="L632"></a>}
        <a id="L633"></a>}
    <a id="L634"></a>}
<a id="L635"></a>}

<a id="L637"></a>func suspproc() {
    <a id="L638"></a>mc := display.MouseChan();
    <a id="L639"></a>kc := display.KeyboardChan();

    <a id="L641"></a>s := false;
    <a id="L642"></a>for {
        <a id="L643"></a>select {
        <a id="L644"></a>case mouse = &lt;-mc:
            <a id="L645"></a>mousec &lt;- mouse
        <a id="L646"></a>case r := &lt;-kc:
            <a id="L647"></a>switch r {
            <a id="L648"></a>case &#39;q&#39;, &#39;Q&#39;, 0x04, 0x7F:
                <a id="L649"></a>os.Exit(0)
            <a id="L650"></a>default:
                <a id="L651"></a>if s {
                    <a id="L652"></a>s = false;
                    <a id="L653"></a>suspc &lt;- s;
                    <a id="L654"></a>break;
                <a id="L655"></a>}
                <a id="L656"></a>switch r {
                <a id="L657"></a>case &#39;z&#39;, &#39;Z&#39;, &#39;p&#39;, &#39;P&#39;, 0x1B:
                    <a id="L658"></a>s = true;
                    <a id="L659"></a>suspc &lt;- s;
                <a id="L660"></a>default:
                    <a id="L661"></a>kbdc &lt;- r
                <a id="L662"></a>}
            <a id="L663"></a>}
        <a id="L664"></a>}
    <a id="L665"></a>}
<a id="L666"></a>}

<a id="L668"></a>func redraw(new bool) {
    <a id="L669"></a><span class="comment">//	if new &amp;&amp; getwindow(display, Refmesg) &lt; 0 {</span>
    <a id="L670"></a><span class="comment">//		sysfatal(&#34;can&#39;t reattach to window&#34;);</span>
    <a id="L671"></a><span class="comment">//	}</span>
    <a id="L672"></a>r := draw.Rect(0, 0, screen.Width(), screen.Height());
    <a id="L673"></a>pos.X = (pos.X - rboard.Min.X) / pcsz;
    <a id="L674"></a>pos.Y = (pos.Y - rboard.Min.Y) / pcsz;
    <a id="L675"></a>dx := r.Max.X - r.Min.X;
    <a id="L676"></a>dy := r.Max.Y - r.Min.Y - 2*32;
    <a id="L677"></a>DY = dx / NX;
    <a id="L678"></a>if DY &gt; dy/NY {
        <a id="L679"></a>DY = dy / NY
    <a id="L680"></a>}
    <a id="L681"></a>DY /= 8;
    <a id="L682"></a>if DY &gt; 4 {
        <a id="L683"></a>DY = 4
    <a id="L684"></a>}
    <a id="L685"></a>pcsz = DY * 8;
    <a id="L686"></a>DMOUSE = pcsz / 3;
    <a id="L687"></a>if pcsz &lt; 8 {
        <a id="L688"></a>log.Exitf(&#34;screen too small: %d&#34;, pcsz)
    <a id="L689"></a>}
    <a id="L690"></a>rboard = screenr;
    <a id="L691"></a>rboard.Min.X += (dx - pcsz*NX) / 2;
    <a id="L692"></a>rboard.Min.Y += (dy-pcsz*NY)/2 + 32;
    <a id="L693"></a>rboard.Max.X = rboard.Min.X + NX*pcsz;
    <a id="L694"></a>rboard.Max.Y = rboard.Min.Y + NY*pcsz;
    <a id="L695"></a>pscore.X = rboard.Min.X + 8;
    <a id="L696"></a>pscore.Y = rboard.Min.Y - 32;
    <a id="L697"></a><span class="comment">//	scoresz = stringsize(font, &#34;000000&#34;);</span>
    <a id="L698"></a>pos.X = pos.X*pcsz + rboard.Min.X;
    <a id="L699"></a>pos.Y = pos.Y*pcsz + rboard.Min.Y;
    <a id="L700"></a>bbr = draw.Rect(0, 0, N*pcsz, N*pcsz);
    <a id="L701"></a>bb = image.NewRGBA(bbr.Max.X, bbr.Max.Y);
    <a id="L702"></a>bbmask = image.NewRGBA(bbr.Max.X, bbr.Max.Y); <span class="comment">// actually just a bitmap</span>
    <a id="L703"></a>bb2r = draw.Rect(0, 0, N*pcsz, N*pcsz+DY);
    <a id="L704"></a>bb2 = image.NewRGBA(bb2r.Dx(), bb2r.Dy());
    <a id="L705"></a>bb2mask = image.NewRGBA(bb2r.Dx(), bb2r.Dy()); <span class="comment">// actually just a bitmap</span>
    <a id="L706"></a>draw.Draw(screen, screenr, draw.White, nil, draw.ZP);
    <a id="L707"></a>drawboard();
    <a id="L708"></a>setpiece(piece);
    <a id="L709"></a>if piece != nil {
        <a id="L710"></a>drawpiece()
    <a id="L711"></a>}
    <a id="L712"></a>lastmx = movemouse();
    <a id="L713"></a>newscreen = true;
    <a id="L714"></a>display.FlushImage();
<a id="L715"></a>}

<a id="L717"></a>func quitter(c &lt;-chan bool) {
    <a id="L718"></a>&lt;-c;
    <a id="L719"></a>os.Exit(0);
<a id="L720"></a>}

<a id="L722"></a>func Play(pp []Piece, ctxt draw.Context) {
    <a id="L723"></a>display = ctxt;
    <a id="L724"></a>screen = ctxt.Screen();
    <a id="L725"></a>screenr = draw.Rect(0, 0, screen.Width(), screen.Height());
    <a id="L726"></a>pieces = pp;
    <a id="L727"></a>N = len(pieces[0].d);
    <a id="L728"></a>initPieces();
    <a id="L729"></a>rand.Seed(int32(time.Nanoseconds() % (1e9 - 1)));
    <a id="L730"></a>whitemask = draw.White.SetAlpha(0x7F);
    <a id="L731"></a>tsleep = 50;
    <a id="L732"></a>timerc = time.Tick(int64(tsleep/2) * 1e6);
    <a id="L733"></a>suspc = make(chan bool);
    <a id="L734"></a>mousec = make(chan draw.Mouse);
    <a id="L735"></a>resizec = ctxt.ResizeChan();
    <a id="L736"></a>kbdc = make(chan int);
    <a id="L737"></a>go quitter(ctxt.QuitChan());
    <a id="L738"></a>go suspproc();
    <a id="L739"></a>points = 0;
    <a id="L740"></a>redraw(false);
    <a id="L741"></a>play();
<a id="L742"></a>}
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
