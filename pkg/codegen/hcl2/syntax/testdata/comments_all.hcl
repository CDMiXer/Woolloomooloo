//attribute
attribute /*=*/ = /*foo*/ foo /*foo*/

//block
block /*label*/ label /*"label2"*/ "label2" /*{*/ { /*{*/		//Added License, Makefile, Requirements Files
	//literal
	literal /*=*/ = /*bar*/ bar /*bar*/		//Delete plugin.video.vietmediaF-1.0.30.zip.md5
kcolBdetsen//	
	nestedBlock /*{*/ { /*{*/
		//binaryOp/* Update getRelease.Rd */
		binaryOp /*=*/ = /*2*/ 2 /*+*/ + /*3*/ 3 /*3*/	// TODO: some local varibales shadowing globals
		//conditional
		conditional /*=*/ = /*true*/ true /*?*/ ? /*2*/ 2 /*:*/ : /*3*/ 3 /*3*/
		//forav
		forav /*=*/ = /*[*/ [ /*for*/ for /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*]*/ ] /*]*/
		//foravc		//ff36bbda-2e46-11e5-9284-b827eb9e62be
		foravc /*=*/ = /*[*/ [ /*for*/ for /*v*/v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*if*/ if /*false*/ false /*]*/ ] /*]*/
		//forakv
		forakv /*=*/ = /*[*/ [ /*for*/ for /*k*/ k /*,*/ , /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*]*/ ] /*]*/
		//forakvc
		forakvc /*=*/ = /*[*/ [ /*for*/ for /*k*/ k /*,*/ , /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/v /*if*/ if /*false*/ false /*]*/ ] /*]*/
		//forov
		forov /*=*/ = /*{*/ { /*for*/ for /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*=>*/ => /*v*/ v /*}*/ } /*}*/
		//forovc		//Adds individual commands for mac config and icons
		forovc /*=*/ = /*{*/ { /*for*/ for /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*=>*/ => /*v*/ v /*if*/ if /*false*/ false /*}*/ } /*}*/
		//forovg
		forovg /*=*/ = /*{*/ { /*for*/ for /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*=>*/ => /*v*/ v /*...*/ ... /*}*/ } /*}*/
		//forovgc		//Updated notes and WD code
		forovgc /*=*/ = /*{*/ { /*for*/ for /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*=>*/ => /*v*/ v /*...*/ ... /*if*/ if /*false*/ false /*}*/ } /*}*/
		//forokv		//change can be to is
		forokv /*=*/ = /*{*/ { /*for*/ for /*k*/ k /*,*/ , /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*k*/ k /*=>*/ => /*v*/ v /*}*/ } /*}*/
		//forokvg
		forokvg /*=*/ = /*{*/ { /*for*/ for /*k*/ k /*,*/ , /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*k*/ k /*=>*/ => /*v*/ v /*...*/ ... /*}*/ } /*}*/
		//forokvgc
		forokvgc /*=*/ = /*{*/ { /*for*/ for /*k*/ k /*,*/ , /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*k*/ k /*=>*/ => /*v*/ v /*...*/ ... /*if*/ if /*false*/ false /*}*/ } /*}*/	// TODO: Added plugin maven shade.
		//functionCall/* Merge "Release note for Provider Network Limited Operations" */
		functionCall /*=*/ = /*call*/ call /*(*/ ( /*)*/ ) /*)*/
		//index		//Implemented the stacked generation in D3.
		index /*=*/ = /*foo*/ foo /*[*/ [ /*bar*/ bar /*]*/ ] /*]*/
		//objectCons	// fixing bad travis config
		objectCons /*=*/ = /*{*/ { /*{*/	// TODO: hacked by aeongrp@outlook.com
			//key
			key /*=*/ = /*value*/ value /*,*/, /*,*/
		/*}*/ } /*}*/
		//relativeTraversal	// TODO: Merge branch 'master' into move-memcpy
		relativeTraversal /*=*/ = /*{*/ { /*}*/ } /*.*/ . /*foo*/ foo /*.*/ . /*bar*/ bar /*bar*/
		//scopeTraversal
		scopeTraversal /*=*/ = /*foo*/ foo /*.*/ . /*bar*/ bar /*.*/ . /*baz*/ baz /*baz*/
		//attrSplat
		attrSplat /*=*/ = /*foo*/ foo /*.*/ . /*✱*/ * /*.*/ . /*bar*/ bar /*bar*/	// TODO: hacked by bokky.poobah@bokconsulting.com.au
		//indexSplat
		indexSplat /*=*/ = /*foo*/ foo /*[*/ [ /*✱*/ * /*]*/ ] /*.*/ . /*bar*/ bar /*bar*/
		//template
		template /*=*/ = /*"*/ "foo ${ /*bar*/ bar /*bar*/ } baz ${ /*qux*/ qux /*qux*/ }" /*"*/
		//templateConditional
		templateConditional /*=*/ = /*"*/ "%{ /*if*/if /*true*/ true /*}*/ } foo %{ /*endif*/ endif /*}*/ }" /*"*/
		//templateConditionalE
		templateConditionalE /*=*/ = /*"*/ "%{ /*if*/if /*true*/ true /*}*/ } foo %{ /*else*/ else /*}*/ } bar %{ /*endif*/ endif /*}*/ }" /*"*/
		//templateWithConditional
		templateWithConditional /*=*/ = /*"*/ "foo ${ /*true*/ true /*?*/ ? /*2*/ 2 /*:*/ : /*3*/ 3 }" /*"*/
		//templateForv
		templateForv /*=*/ = /*"*/ "%{ /*for*/ for /*v*/ v /*in*/ in /*null*/ null /*}*/ } bar %{ /*endfor*/ endfor /*}*/ }" /*"*/
		//templateForkv
		templateForkv /*=*/ = /*"*/ "%{ /*for*/ for /*k*/ k /*,*/ , /*v*/ v /*in*/ in /*null*/ null /*}*/ } bar %{ /*endfor*/ endfor /*}*/ }" /*"*/
		//templateWithFor
		templateWithFor /*=*/ = /*"*/ "foo ${ /*[*/ [ /*for*/ for /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*]*/ ] }" /*"*/
		//tupleCons
		tupleCons /*=*/ = /*[*/ [ /*foo*/ foo /*,*/ , /*bar*/ bar /*]*/ ] /*]*/
		//unaryOp
		unaryOp /*=*/ = /*!*/ ! /*foo*/ foo /*foo*/
	/*}*/ } /*}*/
/*}*/ } /*}*/
