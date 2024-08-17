(0, 954) grammar
  (0, 954) rules
    (0, 18) rule
      (0, 7) name= `grammar`
      (8, 18) alternative
        (12, 17) item
          (12, 17) name= `rules`
    (19, 36) rule
      (19, 24) name= `space`
      (25, 36) alternative
        (29, 35) item
          (30, 34) codepoint= `0020`
    (37, 56) rule
      (37, 44) name= `newline`
      (45, 56) alternative
        (49, 55) item
          (50, 54) codepoint= `000A`
    (57, 89) rule
      (57, 61) name= `name`
      (62, 78) alternative
        (66, 72) item
          (66, 72) name= `letter`
        (73, 77) item
          (73, 77) name= `name`
      (78, 89) alternative
        (82, 88) item
          (82, 88) name= `letter`
    (90, 133) rule
      (90, 96) name= `letter`
      (97, 111) alternative
        (101, 110) item
          (101, 110) range
            (102, 103) codepoint= `a`
            (108, 109) codepoint= `z`
          (110, 110) exclude= ``
      (111, 125) alternative
        (115, 124) item
          (115, 124) range
            (116, 117) codepoint= `A`
            (122, 123) codepoint= `Z`
          (124, 124) exclude= ``
      (125, 133) alternative
        (129, 132) item
          (130, 131) codepoint= `_`
    (134, 174) rule
      (134, 145) name= `indentation`
      (146, 174) alternative
        (150, 155) item
          (150, 155) name= `space`
        (156, 161) item
          (156, 161) name= `space`
        (162, 167) item
          (162, 167) name= `space`
        (168, 173) item
          (168, 173) name= `space`
    (175, 213) rule
      (175, 180) name= `rules`
      (181, 204) alternative
        (185, 189) item
          (185, 189) name= `rule`
        (190, 197) item
          (190, 197) name= `newline`
        (198, 203) item
          (198, 203) name= `rules`
      (204, 213) alternative
        (208, 212) item
          (208, 212) name= `rule`
    (214, 261) rule
      (214, 218) name= `rule`
      (219, 249) alternative
        (223, 227) item
          (223, 227) name= `name`
        (228, 235) item
          (228, 235) name= `newline`
        (236, 248) item
          (236, 248) name= `alternatives`
      (249, 261) alternative
        (253, 260) item
          (253, 260) name= `nothing`
    (262, 309) rule
      (262, 269) name= `nothing`
      (270, 302) alternative
        (274, 285) item
          (274, 285) name= `indentation`
        (286, 289) item
          (287, 288) codepoint= `"`
        (290, 293) item
          (291, 292) codepoint= `"`
        (294, 301) item
          (294, 301) name= `newline`
      (302, 309) alternative
        (306, 308) item
          (307, 307) characters= ``
    (310, 368) rule
      (310, 322) name= `alternatives`
      (323, 352) alternative
        (327, 338) item
          (327, 338) name= `alternative`
        (339, 351) item
          (339, 351) name= `alternatives`
      (352, 368) alternative
        (356, 367) item
          (356, 367) name= `alternative`
    (369, 411) rule
      (369, 380) name= `alternative`
      (381, 411) alternative
        (385, 396) item
          (385, 396) name= `indentation`
        (397, 402) item
          (397, 402) name= `items`
        (403, 410) item
          (403, 410) name= `newline`
    (412, 448) rule
      (412, 417) name= `items`
      (418, 439) alternative
        (422, 426) item
          (422, 426) name= `item`
        (427, 432) item
          (427, 432) name= `space`
        (433, 438) item
          (433, 438) name= `items`
      (439, 448) alternative
        (443, 447) item
          (443, 447) name= `item`
    (449, 475) rule
      (449, 453) name= `item`
      (454, 466) alternative
        (458, 465) item
          (458, 465) name= `literal`
      (466, 475) alternative
        (470, 474) item
          (470, 474) name= `name`
    (476, 539) rule
      (476, 483) name= `literal`
      (484, 498) alternative
        (488, 497) item
          (488, 497) name= `singleton`
      (498, 516) alternative
        (502, 507) item
          (502, 507) name= `range`
        (508, 515) item
          (508, 515) name= `exclude`
      (516, 539) alternative
        (520, 523) item
          (521, 522) codepoint= `"`
        (524, 534) item
          (524, 534) name= `characters`
        (535, 538) item
          (536, 537) codepoint= `"`
    (540, 572) rule
      (540, 549) name= `singleton`
      (550, 572) alternative
        (554, 557) item
          (555, 556) codepoint= `'`
        (558, 567) item
          (558, 567) name= `codepoint`
        (568, 571) item
          (569, 570) codepoint= `'`
    (573, 617) rule
      (573, 582) name= `codepoint`
      (583, 605) alternative
        (587, 604) item
          (587, 604) range
            (588, 592) codepoint= `0020`
            (597, 603) codepoint= `10FFFF`
          (604, 604) exclude= ``
      (605, 617) alternative
        (609, 616) item
          (609, 616) name= `hexcode`
    (618, 695) rule
      (618, 625) name= `hexcode`
      (626, 651) alternative
        (630, 634) item
          (631, 633) characters
            (631, 632) character= `1`
            (632, 633) character= `0`
        (635, 638) item
          (635, 638) name= `hex`
        (639, 642) item
          (639, 642) name= `hex`
        (643, 646) item
          (643, 646) name= `hex`
        (647, 650) item
          (647, 650) name= `hex`
      (651, 675) alternative
        (655, 658) item
          (655, 658) name= `hex`
        (659, 662) item
          (659, 662) name= `hex`
        (663, 666) item
          (663, 666) name= `hex`
        (667, 670) item
          (667, 670) name= `hex`
        (671, 674) item
          (671, 674) name= `hex`
      (675, 695) alternative
        (679, 682) item
          (679, 682) name= `hex`
        (683, 686) item
          (683, 686) name= `hex`
        (687, 690) item
          (687, 690) name= `hex`
        (691, 694) item
          (691, 694) name= `hex`
    (696, 728) rule
      (696, 699) name= `hex`
      (700, 714) alternative
        (704, 713) item
          (704, 713) range
            (705, 706) codepoint= `0`
            (711, 712) codepoint= `9`
          (713, 713) exclude= ``
      (714, 728) alternative
        (718, 727) item
          (718, 727) range
            (719, 720) codepoint= `A`
            (725, 726) codepoint= `F`
          (727, 727) exclude= ``
    (729, 775) rule
      (729, 734) name= `range`
      (735, 775) alternative
        (739, 748) item
          (739, 748) name= `singleton`
        (749, 754) item
          (749, 754) name= `space`
        (755, 758) item
          (756, 757) codepoint= `.`
        (759, 764) item
          (759, 764) name= `space`
        (765, 774) item
          (765, 774) name= `singleton`
    (776, 863) rule
      (776, 783) name= `exclude`
      (784, 822) alternative
        (788, 793) item
          (788, 793) name= `space`
        (794, 797) item
          (795, 796) codepoint= `-`
        (798, 803) item
          (798, 803) name= `space`
        (804, 813) item
          (804, 813) name= `singleton`
        (814, 821) item
          (814, 821) name= `exclude`
      (822, 856) alternative
        (826, 831) item
          (826, 831) name= `space`
        (832, 835) item
          (833, 834) codepoint= `-`
        (836, 841) item
          (836, 841) name= `space`
        (842, 847) item
          (842, 847) name= `range`
        (848, 855) item
          (848, 855) name= `exclude`
      (856, 863) alternative
        (860, 862) item
          (861, 861) characters= ``
    (864, 914) rule
      (864, 874) name= `characters`
      (875, 900) alternative
        (879, 888) item
          (879, 888) name= `character`
        (889, 899) item
          (889, 899) name= `characters`
      (900, 914) alternative
        (904, 913) item
          (904, 913) name= `character`
    (915, 954) rule
      (915, 925) name= `characters`
      (926, 954) alternative
        (930, 953) item
          (930, 947) range
            (931, 935) codepoint= `0020`
            (940, 946) codepoint= `10FFFF`
          (947, 953) exclude
            (951, 952) codepoint= `"`
