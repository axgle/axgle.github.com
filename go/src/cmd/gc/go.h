<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/go.h</title>

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
	<li>Thu Nov 12 15:57:42 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/gc/go.h</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include	&lt;u.h&gt;
#include	&lt;libc.h&gt;
#include	&lt;bio.h&gt;

// avoid &lt;ctype.h&gt;
#undef isblank
#define isblank goisblank

#ifndef	EXTERN
#define	EXTERN	extern
#endif


#define	getc	gcgetc
#define	ungetc	gcungetc
#undef	BUFSIZ

enum
{
	NHUNK		= 50000,
	BUFSIZ		= 8192,
	NSYMB		= 500,
	NHASH		= 1024,
	STRINGSZ	= 200,
	YYMAXDEPTH	= 500,
	MAXALIGN	= 7,
	UINF		= 100,
	HISTSZ		= 10,

	PRIME1		= 3,
	PRIME2		= 10007,
	PRIME3		= 10009,
	PRIME4		= 10037,
	PRIME5		= 10039,
	PRIME6		= 10061,
	PRIME7		= 10067,
	PRIME8		= 10079,
	PRIME9		= 10091,
	PRIME10		= 10093,

	AUNK		= 100,

	// these values are known by runtime
	AMEM		= 0,
	ANOEQ,
	ASTRING,
	AINTER,
	ANILINTER,

	BADWIDTH	= -1000000000
};

/*
 * note this is the representation
 * of the compilers string literals,
 * it is not the runtime representation
 */
typedef	struct	Strlit	Strlit;
struct	Strlit
{
	int32	len;
	char	s[3];	// variable
};

/*
 * note this is the runtime representation
 * of hashmap iterator. it is probably
 * insafe to use it this way, but it puts
 * all the changes in one place.
 * only flag is referenced from go.
 * actual placement does not matter as long
 * as the size is &gt;= actual size.
 */
typedef	struct	Hiter	Hiter;
struct	Hiter
{
	uchar	data[8];		// return val from next
	int32	elemsize;		// size of elements in table */
	int32	changes;		// number of changes observed last time */
	int32	i;			// stack pointer in subtable_state */
	uchar	last[8];		// last hash value returned */
	uchar	h[8];			// the hash table */
	struct
	{
		uchar	sub[8];		// pointer into subtable */
		uchar	start[8];	// pointer into start of subtable */
		uchar	end[8];		// pointer into end of subtable */
		uchar	pad[8];
	} sub[4];
};

enum
{
	Mpscale	= 29,		// safely smaller than bits in a long
	Mpprec	= 16,		// Mpscale*Mpprec is max number of bits
	Mpnorm	= Mpprec - 1,	// significant words in a normalized float
	Mpbase	= 1L &lt;&lt; Mpscale,
	Mpsign	= Mpbase &gt;&gt; 1,
	Mpmask	= Mpbase - 1,
	Mpdebug	= 0,
};

typedef	struct	Mpint	Mpint;
struct	Mpint
{
	long	a[Mpprec];
	uchar	neg;
	uchar	ovf;
};

typedef	struct	Mpflt	Mpflt;
struct	Mpflt
{
	Mpint	val;
	short	exp;
};

typedef	struct	Val	Val;
struct	Val
{
	short	ctype;
	union
	{
		short	reg;		// OREGISTER
		short	bval;		// bool value CTBOOL
		Mpint*	xval;		// int CTINT
		Mpflt*	fval;		// float CTFLT
		Strlit*	sval;		// string CTSTR
	} u;
};

typedef	struct	Sym	Sym;
typedef	struct	Node	Node;
typedef	struct	NodeList	NodeList;
typedef	struct	Type	Type;

struct	Type
{
	uchar	etype;
	uchar	chan;
	uchar	recur;		// to detect loops
	uchar	trecur;		// to detect loops
	uchar	printed;
	uchar	embedded;	// TFIELD embedded type
	uchar	siggen;
	uchar	funarg;
	uchar	copyany;
	uchar	local;		// created in this file
	uchar	deferwidth;
	uchar	broke;

	Node*	nod;		// canonical OTYPE node
	int		lineno;

	// TFUNCT
	uchar	thistuple;
	uchar	outtuple;
	uchar	intuple;
	uchar	outnamed;

	Type*	method;
	Type*	xmethod;

	Sym*	sym;
	int32	vargen;		// unique name for OTYPE/ONAME

	Node*	nname;
	vlong	argwid;

	// most nodes
	Type*	type;
	vlong	width;		// offset in TFIELD, width in all others

	// TFIELD
	Type*	down;		// also used in TMAP
	Strlit*	note;		// literal string annotation

	// TARRAY
	int32	bound;		// negative is dynamic array

	int32	maplineno;	// first use of TFORW as map key
	int32	embedlineno;	// first use of TFORW as embedded type
};
#define	T	((Type*)0)

struct	Node
{
	uchar	op;
	uchar	ullman;		// sethi/ullman number
	uchar	addable;	// type of addressability - 0 is not addressable
	uchar	trecur;		// to detect loops
	uchar	etype;		// op for OASOP, etype for OTYPE, exclam for export
	uchar	class;		// PPARAM, PAUTO, PEXTERN, etc
	uchar	method;		// OCALLMETH name
	uchar	iota;		// OLITERAL made from iota
	uchar	embedded;	// ODCLFIELD embedded type
	uchar	colas;		// OAS resulting from :=
	uchar	diag;		// already printed error about this
	uchar	noescape;	// ONAME never move to heap
	uchar	funcdepth;
	uchar	builtin;	// built-in name, like len or close
	uchar	walkdef;
	uchar	typecheck;
	uchar	local;
	uchar	initorder;
	uchar	dodata;		// compile literal assignment as data statement
	uchar	used;
	uchar	oldref;

	// most nodes
	Node*	left;
	Node*	right;
	Type*	type;
	NodeList*	list;
	NodeList*	rlist;

	// for-body
	NodeList*	ninit;
	Node*	ntest;
	Node*	nincr;
	NodeList*	nbody;

	// if-body
	NodeList*	nelse;

	// cases
	Node*	ncase;

	// func
	Node*	nname;
	Node*	shortname;
	NodeList*	enter;
	NodeList*	exit;
	NodeList*	cvars;	// closure params
	NodeList*	dcl;	// autodcl for this func/closure

	// OLITERAL/OREGISTER
	Val	val;

	// ONAME
	Node*	ntype;
	Node*	defn;
	Node*	pack;	// real package for import . names

	// ONAME func param with PHEAP
	Node*	heapaddr;	// temp holding heap address of param
	Node*	stackparam;	// OPARAM node referring to stack copy of param
	Node*	alloc;	// allocation call

	// ONAME closure param with PPARAMREF
	Node*	outer;	// outer PPARAMREF in nested closure
	Node*	closure;	// ONAME/PHEAP &lt;-&gt; ONAME/PPARAMREF

	Sym*	sym;		// various
	int32	vargen;		// unique name for OTYPE/ONAME
	int32	lineno;
	vlong	xoffset;
	int32	ostk;
};
#define	N	((Node*)0)

struct	NodeList
{
	Node*	n;
	NodeList*	next;
	NodeList*	end;
};

enum
{
	SymExport	= 1&lt;&lt;0,
	SymPackage	= 1&lt;&lt;1,
	SymExported	= 1&lt;&lt;2,
	SymImported	= 1&lt;&lt;3,
	SymUniq		= 1&lt;&lt;4,
	SymSiggen	= 1&lt;&lt;5,
};

struct	Sym
{
	ushort	lexical;
	uchar	flags;
	uchar	sym;		// huffman encoding in object file
	Sym*	link;

	// saved and restored by dcopy
	char*	package;	// package name
	char*	name;		// variable name
	Node*	def;		// definition: ONAME OTYPE OPACK or OLITERAL
	int32	block;		// blocknumber to catch redeclaration
	int32	lastlineno;	// last declaration for diagnostic
};
#define	S	((Sym*)0)

typedef	struct	Iter	Iter;
struct	Iter
{
	int	done;
	Type*	tfunc;
	Type*	t;
	Node**	an;
	Node*	n;
};

typedef	struct	Hist	Hist;
struct	Hist
{
	Hist*	link;
	char*	name;
	int32	line;
	int32	offset;
};
#define	H	((Hist*)0)

enum
{
	OXXX,

	// names
	ONAME,
	ONONAME,
	OTYPE,
	OPACK,
	OLITERAL,

	// exprs
	OADD, OSUB, OOR, OXOR, OADDSTR,
	OADDR,
	OANDAND,
	OAPPENDSTR,
	OARRAY,
	OARRAYBYTESTR, OARRAYRUNESTR,
	OAS, OAS2, OAS2MAPW, OAS2FUNC, OAS2RECV, OAS2MAPR, OAS2DOTTYPE, OASOP,
	OBAD,
	OCALL, OCALLFUNC, OCALLMETH, OCALLINTER,
	OCAP,
	OCLOSE,
	OCLOSED,
	OCLOSURE,
	OCMPIFACE, OCMPSTR,
	OCOMPLIT, OMAPLIT, OSTRUCTLIT, OARRAYLIT,
	OCONV, OCONVNOP, OCONVIFACE, OCONVSLICE,
	ODCL, ODCLFUNC, ODCLFIELD, ODCLCONST, ODCLTYPE,
	ODOT, ODOTPTR, ODOTMETH, ODOTINTER, OXDOT,
	ODOTTYPE,
	OEQ, ONE, OLT, OLE, OGE, OGT,
	OFUNC,
	OIND,
	OINDEX, OINDEXSTR, OINDEXMAP,
	OKEY, OPARAM,
	OLEN,
	OMAKE, OMAKECHAN, OMAKEMAP, OMAKESLICE,
	OHMUL, ORRC, OLRC,	// high-mul and rotate-carry
	OMUL, ODIV, OMOD, OLSH, ORSH, OAND, OANDNOT,
	ONEW,
	ONOT, OCOM, OPLUS, OMINUS,
	OOROR,
	OPANIC, OPANICN, OPRINT, OPRINTN,
	OSEND, OSENDNB,
	OSLICE, OSLICEARR, OSLICESTR,
	ORECV,
	ORUNESTR,
	OSELRECV,
	OIOTA,

	// stmts
	OBLOCK,
	OBREAK,
	OCASE, OXCASE,
	OCONTINUE,
	ODEFER,
	OEMPTY,
	OFALL, OXFALL,
	OFOR,
	OGOTO,
	OIF,
	OLABEL,
	OPROC,
	ORANGE,
	ORETURN,
	OSELECT,
	OSWITCH,
	OTYPECASE,
	OTYPESW,	// l = r.(type)

	// types
	OTCHAN,
	OTMAP,
	OTSTRUCT,
	OTINTER,
	OTFUNC,
	OTARRAY,

	// for back ends
	OCMP, ODEC, OEXTEND, OINC, OREGISTER, OINDREG,

	OEND,
};
enum
{
	Txxx,			// 0

	TINT8,	TUINT8,		// 1
	TINT16,	TUINT16,
	TINT32,	TUINT32,
	TINT64,	TUINT64,
	TINT, TUINT, TUINTPTR,

	TFLOAT32,		// 12
	TFLOAT64,
	TFLOAT,

	TBOOL,			// 15

	TPTR32, TPTR64,		// 16

	TDDD,			// 18
	TFUNC,
	TARRAY,
	T_old_DARRAY,
	TSTRUCT,		// 22
	TCHAN,
	TMAP,
	TINTER,			// 25
	TFORW,
	TFIELD,
	TANY,
	TSTRING,

	// pseudo-types for literals
	TIDEAL,			// 30
	TNIL,
	TBLANK,
	
	// pseudo-type for frame layout
	TFUNCARGS,

	NTYPE,
};
enum
{
	CTxxx,

	CTINT,
	CTFLT,
	CTSTR,
	CTBOOL,
	CTNIL,
};

enum
{
	/* types of channel */
	/* must match ../../pkg/nreflect/type.go:/Chandir */
	Cxxx,
	Crecv = 1&lt;&lt;0,
	Csend = 1&lt;&lt;1,
	Cboth = Crecv | Csend,
};

enum
{
	Pxxx,

	PEXTERN,	// declaration context
	PAUTO,
	PPARAM,
	PPARAMOUT,
	PPARAMREF,	// param passed by reference
	PFUNC,

	PHEAP = 1&lt;&lt;7,
};

enum
{
	Etop = 1&lt;&lt;1,	// evaluated at statement level
	Erv = 1&lt;&lt;2,	// evaluated in value context
	Etype = 1&lt;&lt;3,
	Ecall = 1&lt;&lt;4,	// call-only expressions are ok
	Efnstruct = 1&lt;&lt;5,	// multivalue function returns are ok
	Eiota = 1&lt;&lt;6,		// iota is ok
	Easgn = 1&lt;&lt;7,		// assigning to expression
	Eindir = 1&lt;&lt;8,		// indirecting through expression
	Eaddr = 1&lt;&lt;9,		// taking address of expression
};

#define	BITS	5
#define	NVAR	(BITS*sizeof(uint32)*8)

typedef	struct	Bits	Bits;
struct	Bits
{
	uint32	b[BITS];
};

EXTERN	Bits	zbits;

typedef	struct	Var	Var;
struct	Var
{
	vlong	offset;
	Sym*	sym;
	Sym*	gotype;
	int	width;
	char	name;
	char	etype;
};

EXTERN	Var	var[NVAR];

typedef	struct	Typedef	Typedef;
struct	Typedef
{
	char*	name;
	int	etype;
	int	sameas;
};

extern	Typedef	typedefs[];

typedef	struct	Sig	Sig;
struct Sig
{
	char*	name;
	char*	package;
	Sym*	isym;
	Sym*	tsym;
	Type*	type;
	uint32	hash;
	int32	perm;
	int32	offset;
	Sig*	link;
};

typedef	struct	Io	Io;
struct	Io
{
	char*	infile;
	Biobuf*	bin;
	int32	ilineno;
	int	peekc;
	int	peekc1;	// second peekc for ...
	char*	cp;	// used for content when bin==nil
};

typedef	struct	Dlist	Dlist;
struct	Dlist
{
	Type*	field;
};

typedef	struct	Idir	Idir;
struct Idir
{
	Idir*	link;
	char*	dir;
};

/*
 * argument passing to/from
 * smagic and umagic
 */
typedef	struct	Magic Magic;
struct	Magic
{
	int	w;	// input for both - width
	int	s;	// output for both - shift
	int	bad;	// output for both - unexpected failure

	// magic multiplier for signed literal divisors
	int64	sd;	// input - literal divisor
	int64	sm;	// output - multiplier

	// magic multiplier for unsigned literal divisors
	uint64	ud;	// input - literal divisor
	uint64	um;	// output - multiplier
	int	ua;	// output - adder
};

/*
 * note this is the runtime representation
 * of the compilers arrays.
 *
 * typedef	struct
 * {				// must not move anything
 * 	uchar	array[8];	// pointer to data
 * 	uchar	nel[4];		// number of elements
 * 	uchar	cap[4];		// allocated number of elements
 * } Array;
 */
EXTERN	int	Array_array;	// runtime offsetof(Array,array) - same for String
EXTERN	int	Array_nel;	// runtime offsetof(Array,nel) - same for String
EXTERN	int	Array_cap;	// runtime offsetof(Array,cap)
EXTERN	int	sizeof_Array;	// runtime sizeof(Array)


/*
 * note this is the runtime representation
 * of the compilers strings.
 *
 * typedef	struct
 * {				// must not move anything
 * 	uchar	array[8];	// pointer to data
 * 	uchar	nel[4];		// number of elements
 * } String;
 */
EXTERN	int	sizeof_String;	// runtime sizeof(String)

EXTERN	Dlist	dotlist[10];	// size is max depth of embeddeds

EXTERN	Io	curio;
EXTERN	Io	pushedio;
EXTERN	int32	lexlineno;
EXTERN	int32	lineno;
EXTERN	int32	prevlineno;
EXTERN	char*	pathname;
EXTERN	Hist*	hist;
EXTERN	Hist*	ehist;

EXTERN	char*	infile;
EXTERN	char*	outfile;
EXTERN	char*	package;
EXTERN	Biobuf*	bout;
EXTERN	int	nerrors;
EXTERN	int	nsyntaxerrors;
EXTERN	char	namebuf[NSYMB];
EXTERN	char	lexbuf[NSYMB];
EXTERN	char	debug[256];
EXTERN	Sym*	hash[NHASH];
EXTERN	Sym*	pkgmyname;	// my name for package
EXTERN	Sym*	pkgimportname;	// package name from imported package
EXTERN	int	tptr;		// either TPTR32 or TPTR64
extern	char*	runtimeimport;
extern	char*	unsafeimport;
EXTERN	Idir*	idirs;

EXTERN	Type*	types[NTYPE];
EXTERN	Type*	idealstring;
EXTERN	Type*	idealbool;
EXTERN	uchar	simtype[NTYPE];
EXTERN	uchar	isptr[NTYPE];
EXTERN	uchar	isforw[NTYPE];
EXTERN	uchar	isint[NTYPE];
EXTERN	uchar	isfloat[NTYPE];
EXTERN	uchar	issigned[NTYPE];
EXTERN	uchar	issimple[NTYPE];

EXTERN	uchar	okforeq[NTYPE];
EXTERN	uchar	okforadd[NTYPE];
EXTERN	uchar	okforand[NTYPE];
EXTERN	uchar	okfornone[NTYPE];
EXTERN	uchar	okforcmp[NTYPE];
EXTERN	uchar	okforbool[NTYPE];
EXTERN	uchar	okforcap[NTYPE];
EXTERN	uchar	okforlen[NTYPE];
EXTERN	uchar	okforarith[NTYPE];
EXTERN	uchar*	okfor[OEND];
EXTERN	uchar	iscmp[OEND];

EXTERN	Mpint*	minintval[NTYPE];
EXTERN	Mpint*	maxintval[NTYPE];
EXTERN	Mpflt*	minfltval[NTYPE];
EXTERN	Mpflt*	maxfltval[NTYPE];

EXTERN	NodeList*	xtop;
EXTERN	NodeList*	externdcl;
EXTERN	NodeList*	closures;
EXTERN	NodeList*	exportlist;
EXTERN	NodeList*	typelist;
EXTERN	int	dclcontext;		// PEXTERN/PAUTO
EXTERN	int	incannedimport;
EXTERN	int	statuniqgen;		// name generator for static temps
EXTERN	int	loophack;

EXTERN	uint32	iota;
EXTERN	NodeList*	lastconst;
EXTERN	Node*	lasttype;
EXTERN	int32	maxarg;
EXTERN	int32	stksize;		// stack size for current frame
EXTERN	int32	blockgen;		// max block number
EXTERN	int32	block;			// current block number
EXTERN	int	hasdefer;		// flag that curfn has defer statetment

EXTERN	Node*	curfn;

EXTERN	int	maxround;
EXTERN	int	widthptr;

EXTERN	Node*	typesw;
EXTERN	Node*	nblank;

EXTERN	char*	structpkg;
extern	int	thechar;
extern	char*	thestring;
EXTERN	char*	hunk;
EXTERN	int32	nhunk;
EXTERN	int32	thunk;

EXTERN	int	exporting;
EXTERN	int	noargnames;

EXTERN	int	funcdepth;
EXTERN	int	typecheckok;

/*
 *	y.tab.c
 */
int	yyparse(void);

/*
 *	lex.c
 */
void	addidir(char*);
void	importfile(Val*, int line);
void	cannedimports(char*, char*);
void	unimportfile();
int32	yylex(void);
void	yyoptsemi(int);
void	typeinit(void);
void	lexinit(void);
char*	lexname(int);
int32	getr(void);
int	escchar(int, int*, vlong*);
int	getc(void);
void	ungetc(int);
void	mkpackage(char*);

/*
 *	mparith1.c
 */
int	mpcmpfixflt(Mpint *a, Mpflt *b);
int	mpcmpfltfix(Mpflt *a, Mpint *b);
int	mpcmpfixfix(Mpint *a, Mpint *b);
int	mpcmpfixc(Mpint *b, vlong c);
int	mpcmpfltflt(Mpflt *a, Mpflt *b);
int	mpcmpfltc(Mpflt *b, double c);
void	mpsubfixfix(Mpint *a, Mpint *b);
void	mpsubfltflt(Mpflt *a, Mpflt *b);
void	mpaddcfix(Mpint *a, vlong c);
void	mpaddcflt(Mpflt *a, double c);
void	mpmulcfix(Mpint *a, vlong c);
void	mpmulcflt(Mpflt *a, double c);
void	mpdivfixfix(Mpint *a, Mpint *b);
void	mpmodfixfix(Mpint *a, Mpint *b);
void	mpatofix(Mpint *a, char *s);
void	mpatoflt(Mpflt *a, char *s);
int	mpmovefltfix(Mpint *a, Mpflt *b);
void	mpmovefixflt(Mpflt *a, Mpint *b);
int	Bconv(Fmt*);

/*
 *	mparith2.c
 */
void	mpmovefixfix(Mpint *a, Mpint *b);
void	mpmovecfix(Mpint *a, vlong v);
int	mptestfix(Mpint *a);
void	mpaddfixfix(Mpint *a, Mpint *b);
void	mpmulfixfix(Mpint *a, Mpint *b);
void	mpmulfract(Mpint *a, Mpint *b);
void	mpdivmodfixfix(Mpint *q, Mpint *r, Mpint *n, Mpint *d);
void	mpdivfract(Mpint *a, Mpint *b);
void	mpnegfix(Mpint *a);
void	mpandfixfix(Mpint *a, Mpint *b);
void	mpandnotfixfix(Mpint *a, Mpint *b);
void	mplshfixfix(Mpint *a, Mpint *b);
void	mporfixfix(Mpint *a, Mpint *b);
void	mprshfixfix(Mpint *a, Mpint *b);
void	mpxorfixfix(Mpint *a, Mpint *b);
void	mpcomfix(Mpint *a);
vlong	mpgetfix(Mpint *a);
void	mpshiftfix(Mpint *a, int s);

/*
 *	mparith3.c
 */
int	sigfig(Mpflt *a);
void	mpnorm(Mpflt *a);
void	mpmovefltflt(Mpflt *a, Mpflt *b);
void	mpmovecflt(Mpflt *a, double f);
int	mptestflt(Mpflt *a);
void	mpaddfltflt(Mpflt *a, Mpflt *b);
void	mpmulfltflt(Mpflt *a, Mpflt *b);
void	mpdivfltflt(Mpflt *a, Mpflt *b);
void	mpnegflt(Mpflt *a);
double	mpgetflt(Mpflt *a);
int	Fconv(Fmt*);

/*
 *	subr.c
 */
void*	mal(int32);
void*	remal(void*, int32, int32);
void	errorexit(void);
uint32	stringhash(char*);
Sym*	lookup(char*);
Sym*	pkglookup(char*, char*);
Sym*	restrictlookup(char*, char*);
void	importdot(Sym*, Node*);
void	yyerror(char*, ...);
void	yyerrorl(int, char*, ...);
void	flusherrors(void);
int	parserline(void);
void	warn(char*, ...);
void	fatal(char*, ...);
void	linehist(char*, int32, int);
int32	setlineno(Node*);
Node*	nod(int, Node*, Node*);
Node*	nodlit(Val);
Type*	typ(int);
int	algtype(Type*);
void	dodump(Node*, int);
void	dump(char*, Node*);
void	dumplist(char*, NodeList*);
Type*	aindex(Node*, Type*);
int	isnil(Node*);
int	isptrto(Type*, int);
int	istype(Type*, int);
int	isfixedarray(Type*);
int	isslice(Type*);
int	isinter(Type*);
int	isnilinter(Type*);
int	isddd(Type*);
int	isideal(Type*);
int	isblank(Node*);
Type*	maptype(Type*, Type*);
Type*	methtype(Type*);
Node*	typename(Type*);
int	eqtype(Type*, Type*);
int	cvttype(Type*, Type*);
int	eqtypenoname(Type*, Type*);
void	argtype(Node*, Type*);
int	eqargs(Type*, Type*);
uint32	typehash(Type*);
void	frame(int);
Node*	nodintconst(int64);
void	nodconst(Node*, Type*, int64);
Node*	nodnil(void);
Node*	nodbool(int);
void	ullmancalc(Node*);
void	badtype(int, Type*, Type*);
Type*	ptrto(Type*);
NodeList*	cleanidlist(NodeList*);
Node*	syslook(char*, int);
Node*	treecopy(Node*);
NodeList*	listtreecopy(NodeList*);
int	isselect(Node*);
Node*	staticname(Type*);
int	iscomposite(Type*);
Node*	callnew(Type*);
Node*	saferef(Node*, NodeList**);
Node*	safeval(Node*, NodeList**);
int	is64(Type*);
int	noconv(Type*, Type*);
NodeList*	list1(Node*);
NodeList*	list(NodeList*, Node*);
NodeList*	concat(NodeList*, NodeList*);
int		count(NodeList*);
Node*	liststmt(NodeList*);

Type**	getthis(Type*);
Type**	getoutarg(Type*);
Type**	getinarg(Type*);

Type*	getthisx(Type*);
Type*	getoutargx(Type*);
Type*	getinargx(Type*);

Type*	structfirst(Iter*, Type**);
Type*	structnext(Iter*);
Type*	funcfirst(Iter*, Type*);
Type*	funcnext(Iter*);

int	brcom(int);
int	brrev(int);
void	setmaxarg(Type*);
int	dotoffset(Node*, int*, Node**);
void	tempname(Node*, Type*);

int	Econv(Fmt*);
int	Jconv(Fmt*);
int	Lconv(Fmt*);
int	Oconv(Fmt*);
int	Sconv(Fmt*);
int	Tconv(Fmt*);
int	Nconv(Fmt*);
void	exprfmt(Fmt*, Node*, int);
int	Wconv(Fmt*);
int	Zconv(Fmt*);

int	lookdot0(Sym*, Type*, Type**);
int	adddot1(Sym*, Type*, int, Type**);
Node*	adddot(Node*);
void	expandmeth(Sym*, Type*);
void	genwrapper(Type*, Type*, Sym*);

int	simsimtype(Type*);

int	powtwo(Node*);
Type*	tounsigned(Type*);
void	smagic(Magic*);
void	umagic(Magic*);

void	redeclare(Sym*, char*);
Sym*	ngotype(Node*);

/*
 *	dcl.c
 */
void	declare(Node*, int);
Type*	dodcltype(Type*);
void	updatetype(Type*, Type*);
void	defaultlit(Node**, Type*);
void	defaultlit2(Node**, Node**, int);
int	structcount(Type*);
void	addmethod(Sym*, Type*, int);
Node*	methodname(Node*, Type*);
Node*	methodname1(Node*, Node*);
Sym*	methodsym(Sym*, Type*);
Type*	functype(Node*, NodeList*, NodeList*);
char*	thistypenam(Node*);
void	funcnam(Type*, char*);
Node*	renameinit(Node*);
void	funchdr(Node*);
void	funcbody(Node*);
Node*	typenod(Type*);
Type*	dostruct(NodeList*, int);
Type**	stotype(NodeList*, int, Type**);
Type*	sortinter(Type*);
void	markdcl(void);
void	popdcl(void);
void	poptodcl(void);
void	dumpdcl(char*);
void	markdclstack(void);
void	testdclstack(void);
Sym*	pushdcl(Sym*);
void	addvar(Node*, Type*, int);
void	addtyp(Type*, int);
void	addconst(Node*, Node*, int);
Node*	fakethis(void);
int	isifacemethod(Type*);
Node*	dclname(Sym*);
Node*	newname(Sym*);
Node*	oldname(Sym*);
Type*	newtype(Sym*);
Type*	oldtype(Sym*);
void	fninit(NodeList*);
Node*	nametodcl(Node*, Type*);
NodeList*	checkarglist(NodeList*);
void	checkwidth(Type*);
void	defercheckwidth(void);
void	resumecheckwidth(void);
Node*	embedded(Sym*);
NodeList*	variter(NodeList*, Node*, NodeList*);
NodeList*	constiter(NodeList*, Node*, NodeList*);

Node*	unsafenmagic(Node*, NodeList*);
void	dclchecks(void);
void	funccompile(Node*);

Node*	typedcl0(Sym*);
Node*	typedcl1(Node*, Node*, int);
void	typedcl2(Type*, Type*);

/*
 * closure.c
 */
void	closurehdr(Node*);
Node*	closurebody(NodeList*);
void	typecheckclosure(Node*);
Node*	walkclosure(Node*, NodeList**);


/*
 * sinit.c
 */

NodeList*	initfix(NodeList*);

/*
 *	export.c
 */
void	autoexport(Node*, int);
int	exportname(char*);
void	exportsym(Node*);
void	packagesym(Node*);
void	dumpe(Sym*);
void	dumpexport(void);
void	dumpexporttype(Sym*);
void	dumpexportvar(Sym*);
void	dumpexportconst(Sym*);
void	importconst(Sym *s, Type *t, Node *v);
void	importmethod(Sym *s, Type *t);
void	importtype(Type *s, Type *t);
void	importvar(Sym *s, Type *t, int ctxt);
void	checkimports(void);
Type*	pkgtype(Sym*);
Sym*	importsym(Sym*, int);

/*
 *	walk.c
 */
void	walk(Node*);
void	walkstmt(Node**);
void	walkstmtlist(NodeList*);
void	walkexprlist(NodeList*, NodeList**);
void	walkconv(Node**, NodeList**);
void	walkdottype(Node*, NodeList**);
void	walkas(Node*);
void	walkswitch(Node*);
void	walkrange(Node*);
void	walkselect(Node*);
void	walkdot(Node*, NodeList**);
void	walkexpr(Node**, NodeList**);
Node*	mkcall(char*, Type*, NodeList**, ...);
Node*	mkcall1(Node*, Type*, NodeList**, ...);
Node*	chanfn(char*, int, Type*);
Node*	ascompatee1(int, Node*, Node*, NodeList**);
NodeList*	ascompatee(int, NodeList*, NodeList*, NodeList**);
NodeList*	ascompatet(int, NodeList*, Type**, int, NodeList**);
NodeList*	ascompatte(int, Type**, NodeList*, int, NodeList**);
Node*	mapop(Node*, NodeList**);
Type*	fixchan(Type*);
Node*	ifacecvt(Type*, Node*, int, NodeList**);
int	ifaceas(Type*, Type*, int);
int	ifaceas1(Type*, Type*, int);
void	ifacecheck(Type*, Type*, int, int);
void	runifacechecks(void);
Node*	convas(Node*, NodeList**);
Node*	colas(NodeList*, NodeList*);
void	colasdefn(NodeList*, Node*);
NodeList*	reorder1(NodeList*);
NodeList*	reorder3(NodeList*);
NodeList*	reorder4(NodeList*);
int	vmatch1(Node*, Node*);
void	anylit(Node*, Node*, NodeList**);
int	oaslit(Node*, NodeList**);
void	heapmoves(void);
void	walkdeflist(NodeList*);
void	walkdef(Node*);
void	typechecklist(NodeList*, int);
void	typecheckswitch(Node*);
void	typecheckselect(Node*);
void	typecheckrange(Node*);
Node*	typecheckconv(Node*, Node*, Type*, int);
int	checkconv(Type*, Type*, int, int*, int*);
Node*	typecheck(Node**, int);

/*
 *	const.c
 */
void	convlit1(Node**, Type*, int);
void	convlit(Node**, Type*);
void	evconst(Node*);
int	cmpslit(Node *l, Node *r);
int	smallintconst(Node*);
long	nonnegconst(Node*);
int	consttype(Node*);
int	isconst(Node*, int);
Mpflt*	truncfltlit(Mpflt*, Type*);
void	convconst(Node*, Type*, Val*);

/*
 *	align.c
 */
uint32	rnd(uint32, uint32);
void	dowidth(Type*);
int	argsize(Type*);

/*
 *	bits.c
 */
Bits	bor(Bits, Bits);
Bits	band(Bits, Bits);
Bits	bnot(Bits);
int	bany(Bits*);
int	bnum(Bits);
Bits	blsh(uint);
int	beq(Bits, Bits);
int	bset(Bits, uint);
int	Qconv(Fmt *fp);
int	bitno(int32);

/*
 *	gen.c
 */
typedef	struct	Prog	Prog;
#define	P	((Prog*)0)

typedef	struct	Label Label;
struct	Label
{
	uchar	op;		// OGOTO/OLABEL
	Sym*	sym;
	Node*	stmt;
	Prog*	label;		// pointer to code
	Prog*	breakpc;	// pointer to code
	Prog*	continpc;	// pointer to code
	Label*	link;
};
#define	L	((Label*)0)

EXTERN	Label*	labellist;

typedef	struct	Plist	Plist;
struct	Plist
{
	Node*	name;
	Prog*	firstpc;
	int	recur;
	Plist*	link;
};

EXTERN	Plist*	plist;
EXTERN	Plist*	plast;

EXTERN	Prog*	continpc;
EXTERN	Prog*	breakpc;
EXTERN	Prog*	pc;
EXTERN	Prog*	firstpc;

EXTERN	int	yylast;
EXTERN	int	yynext;
EXTERN	int	yysemi;

void	allocparams(void);
void	cgen_as(Node *nl, Node *nr);
void	cgen_callmeth(Node *n, int proc);
void	cgen_dcl(Node *n);
void	cgen_proc(Node *n, int proc);
void	checklabels(void);
void	gen(Node *n);
void	genlist(NodeList *l);
void	newlab(int op, Sym *s, Node*);
Node*	sysfunc(char *name);
Plist*	newplist(void);

/*
 *	obj.c
 */
void	Bputdot(Biobuf *b);
void	dumpglobls(void);
void	dumpobj(void);
void	ieeedtod(uint64 *ieee, double native);
void	outhist(Biobuf *b);

/*
 *	arch-specific gen.c/gsubr.c/obj.c
 */
void	betypeinit(void);
vlong	convvtox(vlong, int);
void	compile(Node*);
void	proglist(void);
int	optopop(int);
void	dumpobj(void);
void	dowidth(Type*);
void	argspace(int32);
Node*	nodarg(Type*, int);
Type*	deep(Type*);
Type*	shallow(Type*);
Prog*	gjmp(Prog*);
void	patch(Prog*, Prog*);
void	bgen(Node *n, int true, Prog *to);
void	cgen_asop(Node *n);
void	cgen_call(Node *n, int proc);
void	cgen_callinter(Node *n, Node *res, int proc);
void	cgen_ret(Node *n);
int	isfat(Type*);
void	clearfat(Node *n);
void	cgen(Node*, Node*);
struct Prog;
void	gused(Node*);
void	gdata(Node*, Node*, int);
void	gdatastring(Node*, Strlit*);
void	dumptypestructs(void);
void	dumpfuncs(void);
void	dumpdata(void);
void	ggloblnod(Node *nam, int32 width);
void	ggloblsym(Sym *s, int32 width, int dupok);
void	zfile(Biobuf *b, char *p, int n);
void	zhist(Biobuf *b, int line, vlong offset);
void	zname(Biobuf *b, Sym *s, int t);
void	nopout(Prog*);
int	dstringptr(Sym *s, int off, char *str);
int	dgostringptr(Sym*, int off, char *str);
int	dgostrlitptr(Sym*, int off, Strlit*);
int	dsymptr(Sym *s, int off, Sym *x, int xoff);
int	duint8(Sym *s, int off, uint8 v);
int	duint16(Sym *s, int off, uint16 v);
int	duint32(Sym *s, int off, uint32 v);
int	duint64(Sym *s, int off, uint64 v);
int	duintptr(Sym *s, int off, uint64 v);
int	duintxx(Sym *s, int off, uint64 v, int wid);
void	genembedtramp(Type*, Type*, Sym*);
int	gen_as_init(Node*);

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
