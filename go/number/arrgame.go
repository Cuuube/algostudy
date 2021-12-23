package main

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
)

func main() {
	t1 := []int{3, 4, 5, 1, 6, 7}
	t2 := []int{1, 2, 3, 4, 5}
	t3 := []int{1, 1, 1, 2, 3, 4}

	fmt.Println(numsGame(t1))
	fmt.Println(numsGame(t2))
	fmt.Println(numsGame(t3))

	testCase := `[225,811,502,88,874,55,280,459,743,903,347,428,144,629,511,452,385,561,201,299,580,335,542,173,18,523,401,787,934,818,636,738,658,630,475,577,463,535,660,607,628,973,459,993,933,478,589,711,626,760,599,361,761,874,697,72,432,258,224,995,584,666,212,36,98,291,114,231,873,776,819,710,61,196,924,919,578,744,57,710,437,936,175,307,389,185,634,869,578,645,9,12,718,80,253,258,803,373,161,122,867,957,31,816,467,926,494,20,546,416,479,15,414,832,293,941,568,767,593,857,533,304,925,185,313,247,77,264,405,321,638,910,860,68,910,443,785,779,500,573,804,120,164,477,806,267,926,405,219,667,117,314,651,269,410,129,695,854,879,556,492,687,90,885,668,13,104,150,578,87,63,405,448,477,816,991,609,465,32,621,355,153,231,320,33,329,270,449,379,636,33,478,526,541,395,769,847,255,86,188,15,997,601,142,815,223,381,127,254,422,483,138,207,501,560,115,27,287,458,386,962,131,397,818,721,481,250,806,459,527,425,887,969,755,321,224,610,357,888,340,797,695,406,373,244,405,591,400,488,855,481,988,105,655,455,421,143,708,746,564,506,206,831,59,825,726,990,863,467,168,585,687,569,840,559,721,479,573,885,751,882,560,177,194,292,330,891,931,919,97,479,129,932,707,790,537,201,835,410,520,97,369,846,272,188,37,129,895,440,531,665,797,209,576,466,926,411,846,603,670,992,956,489,440,837,927,734,778,220,82,644,193,699,509,276,794,388,126,83,263,93,554,783,780,26,716,63,672,972,679,876,67,437,614,32,958,964,110,902,500,720,33,101,798,886,298,248,915,846,85,924,968,308,717,441,234,598,241,342,88,912,192,220,536,331,957,366,737,960,303,397,142,222,449,224,127,322,889,821,291,205,738,619,386,443,225,959,799,800,503,360,796,849,662,597,456,395,512,837,918,541,154,781,862,44,717,554,498,581,953,630,132,494,35,522,517,621,150,600,531,868,292,489,453,590,20,716,37,803,365,91,682,830,854,500,629,530,767,603,918,726,123,723,116,546,422,795,356,438,277,542,40,985,121,240,494,4,488,847,800,610,257,386,334,609,287,16,363,891,684,706,367,825,758,273,952,328,788,910,797,905,303,180,472,362,142,616,836,203,884,470,818,502,185,470,829,342,611,77,92,50,838,475,832,271,119,441,293,291,360,31,843,109,954,719,617,166,249,785,806,736,98,168,913,974,728,927,968,274,549,973,503,254,964,296,547,908,773,695,259,818,257,776,870,201,9,398,428,492,733,705,404,249,330,317,361,22,74,862,869,356,978,518,710,365,214,294,650,9,120,940,621,352,409,939,99,296,779,488,951,856,931,848,427,75,316,329,178,484,262,14,159,926,164,531,618,794,953,388,234,403,552,922,897,835,405,975,157,631,958,176,989,466,107,934,408,321,467,836,594,555,503,784,723,772,795,183,576,366,745,622,642,661,781,554,40,92,715,964,960,697,259,177,417,23,810,971,336,633,939,596,618,795,254,206,209,995,38,738,266,523,31,354,100,15,384,993,392,165,270,17,273,397,926,94,124,544,330,315,607,508,702,497,557,969,911,920,962,918,468,382,532,387,884,572,726,332,877,609,603,339,530,513,987,676,428,63,52,74,321,790,885,901,308,856,130,390,475,842,616,29,937,65,629,516,533,609,65,102,344,239,951,344,913,717,173,148,579,234,897,437,116,908,789,903,199,550,982,219,647,980,288,982,685,580,836,22,388,8,4,997,156,34,959,741,559,35,28,635,656,482,672,509,280,882,255,953,692,989,671,915,958,400,839,117,394,398,587,894,115,450,131,956,855,97,962,639,217,935,358,172,962,845,202,426,274,306,65,930,39,80,306,631,211,194,495,692,477,620,637,939,930,13,346,229,159,283,767,246,418,580,393,424,77,992,132,56,195,185,124,364,849,961,54,396,394,898,24,738,406,309,452,207,335,164,777,86,492,770,703,844,6,586,509,153,445,393,195,316,334,846,540,875,102,635,249,597,442,967,863,556,189,160,863,166,832,645,791,704,909,549,715,396,81,511,854,541,353,122,915,390,949,385,119,767,507,301,246,812,667,137,212,702,51,638,435,188,289,148,776,161,960,180,366,753,501,530,907,938,327,278,350,886,765,530,912,455,42,49,494,328,350,332,3,849,993,530,876,866,953,719,989,618,589,477,198,320,354,328,250,593,625,295,821,118,136,94,389,847,26,967,279,905,615,763,190,902,309,340,432,973,51,672,935,298,205,713,888,575,266,715,1000,441,272,270,816,951,307,73,757,461,827,308,560,12,120,611,405,945,541,322,408,945,414,408,457,758,705,653,37,102,806,781,156,872,813,920,135,53,656,869,682,226,920,466,927,239,187,202,176,325,232,587,514,893,776,399,650,416,26,320,15,612,293,494,536,280,203,982,895,315,110,747,908,524,259,673,187,120,584,235,714,251,756,801,676,223,619,93,967,822,762,367,466,27,29,645,994,766,176,462,291,159,501,617,231,653,509,330,87,974,10,516,378,522,904,652,904,494,614,141,523,535,958,840,941,787,587,447,57,320,634,869,443,351,430,669,798,2,823,165,896,77,245,604,49,880,586,285,47,330,293,943,898,59,112,727,684,49,19,172,608,13,90,427,270,901,13,319,226,757,189,230,143,51,687,143,793,265,733,544,718,213,372,82,413,453,381,344,656,877,389,210,882,390,93,995,703,526,298,382,508,768,475,910,850,470,19,303,247,47,303,641,400,603,363,871,773,792,780,754,124,359,914,454,899,137,883,991,99,854,997,749,58,200,483,17,302,443,74,865,51,909,450,975,405,90,294,973,677,368,816,100,529,509,277,626,936,776,264,857,190,469,936,985,856,761,413,297,370,204,393,583,114,695,478,727,455,511,966,846,793,377,13,843,66,9,595,367,115,44,493,858,184,158,158,316,399,851,553,174,60,300,569,298,331,448,397,496,767,467,281,587,239,363,511,596,57,416,279,390,706,393,237,392,301,390,430,987,481,412,806,619,967,21,335,627,915,862,683,801,21,721,330,838,461,425,672,17,321,407,477,211,541,395,888,983,687,911,469,98,409,12,611,752,659,136,727,532,660,192,586,536,654,767,586,638,374,792,92,229,831,781,789,706,886,946,756,110,580,616,925,811,126,272,726,974,725,915,845,42,189,778,597,146,1000,218,33,258,693,35,73,889,20,819,476,475,371,409,234,262,89,527,777,616,823,230,375,395,757,742,495,200,194,176,622,650,279,607,992,847,739,872,601,998,395,926,564,564,926,818,177,613,967,18,141,663,12,943,9,742,537,242,384,671,672,884,94,789,868,882,197,134,600,521,54,943,468,14,358,554,260,863,862,638,639,670,564,203,796,163,529,403,125,672,518,950,962,157,601,965,536,749,274,233,729,607,196,663,731,63,489,853,17,268,658,641,776,398,21,665,296,477,421,766,427,408,105,82,874,131,980,687,229,745,640,756,64,840,492,523,929,861,78,430,987,715,782,632,666,99,414,471,301,454,436,718,879,786,716,894,739,22,797,168,387,904,979,658,562,562,533,213,431,522,629,996,117,521,917,84,274,525,69,177,419,547,483,906,644,432,795,13,640,231,101,346,971,109,500,218,225,115,628,797,279,258,557,352,546,505,888,287,753,597,567,778,897,950,896,385,678,55,131,787,789,562,198,325,300,280,382,773,974,821,251,849,286,469,393,29,265,509,425,535,33,543,715,740,211,327,928,470,794,396,527,654,94,314,759,7,329,700,775,347,823,175,441,650,451,299,970,567,1,199,108,720,994,368,715,176,887,271,179,901,248,216,889,885,183,601,724,332,738,438,132,860,143,149,235,848,938,435,387,356,390,753,918,622,756,189,359,858,12,908,865,717,538,460,882,118,487,737,217,316,772,540,20,650,729,621,832,873,324,42,379,146,22,454,990,809,802,39,206,352,992,32,182,105,455,455,898,547,958,49,738,364,207,567,77,789,404,382,416,121,768,445,45,753,539,197,766,436,499,453,476,514,471,320,285,844,686,828,686,35,211,658,377,125,51,75,998,714,212,851,578,376,780,879,144,447,673,413,594,535,686,129,382,780,773,858,912,4,349,678,402,995,357,589,199,951,367,841,582,907,638,967,651,410,30,164,718,27,530,956,640,398,722,226,637,896,170,351,435,750,280,994,542,754,156,988,93,98,133,9,588,270,163,449,810,742,227,355,74,205,923,315,511,232,29,290,20,953,883,731,662,320,636,297,260,55,688,797,37,563,371,979,462,623,473,334,253,424,783,655,521,900,568,936,380,544,595,766,77,798,161,834,560,452,211,922,707,193,963,750,97,580,589,998,206,195,948,143,383,96,10,16,317,712,787,657,214,972,354,692,183,335,297,392,178,571,122,790,796,240,754,191,524,576,840,589,295,572,793,66,306,226,473,896,565,820,879,585,400,495,683,294,439,77,977,343,159,534,23,622,406,156,415,251,527,262,778,28,622,805,630,824,226,452,340,192,217,613,1000,526,993,629,222,368,664,707,1000,126,216,652,305,207,200,355,681,206,817,558,377,284,774,41,202,177,116,535,121,577,458,391,823,349,719,270,294,760,646,137,117,763,304,798,238,90,906,251,680,128,150,845,454,853,373,359,323,500,818,438,518,729,309,240,276,949,920,445,588,457,377,570,476,786,959,273,968,202,252,297,740,110,297,329,609,198,766,609,140,482,5,396,365,434,731,273,893,311,804,627,460,458,749,3,819,805,307,455,837,591,721,511,461,668,581,796,923,697,836,744,177,927,215,611,986,636,755,834,663,847,973,755,32,330,818,473,569,909,455,132,888,120,960,418,253,179,397,164,180,328,692,527,571,207,288,562,9,228,717,599,390,865,767,59,663,538,554,100,901,16,607,923,968,506,184,158,289,681,656,639,4,513,570,896,692,225,548,843,238,547,7,144,150,835,667,134,492,367,342,202,252,340,945,720,776,197,607,537,298,709,1000,309,818,408,682,76,262,522,50,157,179,252,249,257,191,579,212,717,318,694,763,825,280,406,532,615,481,569,214,332,891,658,771,579,389,580,879,704,544,81,379,666,893,460,922,814,58,602,258,971,105,506,966,139,176,169,15,410,680,600,662,444,318,175,73,964,563,733,279,174,984,452,382,725,729,555,450,445,192,273,290,620,180,976,980,800,925,685,137,162,704,335,413,153,503,720,325,46,992,436,820,922,214,6,432,201,455,522,568,814,868,507,68,466,114,868,373,926,251,44,38,829,943,851,700,745,679,489,982,255,918,691,964,329,849,845,732,349,536,126,189,15,174,770,916,10,479,464,989,530,598,83,833,38,352,384,74,510,984,598,80,403,876,437,731,187,814,129,285,335,631,607,521,658,506,381,434,967,11,34,884,282,40,835,554,722,916,705,540,20,712,236,627,153,347,746,497,845,585,946,510,67,828,697,512,409,250,460,281,598,858,789,797,508,502,26,669,616,516,151,166,192,132,771,684,752,748,469,509,151,122,146,580,410,139,11,616,117,234,436,323,963,181,575,185,730,803,211,473,53,570,934,336,801,790,946,961,205,383,572,434,528,108,617,628,30,482,671,422,660,701,318,593,233,969,910,739,889,872,369,346,985,87,981,819,944,346,993,941,996,423,660,41,891,95,728,194,805,891,611,957,229,975,860,502,767,497,275,42,999,560,699,679,805,616,267,363,344,965,578,743,494,736,298,385,651,3,289,945,673,589,767,295,700,756,387,791,283,825,754,914,411,496,559,964,455,805,610,2,469,846,734,596,803,151,580,235,800,306,462,591,9,885,509,249,215,745,313,5,978,658,161,193,618,355,979,923,813,151,737,371,145,397,592,579,158,809,621,547,941,466,738,697,748,20,633,228,586,660,838,329,315,299,208,619,44,872,821,370,934,86,732,350,124,580,587,442,948,960,226,188,230,823,559,145,540,921,777,248,431,796,642,894,593,251,308,884,653,628,665,798,779,99,131,885,323,710,751,477,701,835,778,113,621,404,26,378,622,793,557,436,101,754,434,564,37,193,708,115,415,293,423,509,193,886,359,792,988,655,528,984,312,974,277,673,606,434,735,713,668,16,270,911,937,366,314,224,882,729,584,553,313,364,897,158,175,368,609,107,517,221,504,416,783,823,169,674,540,430,604,429,448,988,116,292,89,442,235,632,722,576,141,813,360,887,975,436,584,892,720,967,426,158,812,567,328,526,358,723,684,927,966,271,843,847,235,517,545,957,289,87,654,321,24,450,967,368,381,999,873,186,640,356,2,250,910,610,803,945,568,310,921,381,672,286,866,736,140,655,211,352,575,241,565,667,207,849,192,776,74,306,25,603,709,975,886,214,729,450,72,954,221,781,907,40,837,936,290,455,309,546,509,80,616,44,11,56,30,320,979,61,226,299,685,248,258,637,706,461,511,267,54,866,243,106,487,973,55,466,494,805,356,750,514,869,271,695,964,157,443,221,965,924,415,468,82,86,101,692,289,4,857,991,510,427,218,676,935,61,492,824,78,981,130,301,979,657,826,735,296,52,77,228,21,2,84,583,997,405,5,508,422,83,310,763,184,467,720,680,984,84,49,691,930,192,196,772,820,753,197,374,331,433,744,194,653,31,507,210,833,277,77,906,972,698,789,151,326,934,572,502,323,317,408,74,378,811,406,900,714,499,17,145,822,405,767,170,20,991,979,685,291,254,380,79,141,593,972,966,161,450,535,102,294,930,927,399,718,345,67,643,539,20,262,924,188,364,818,12,867,870,612,29,675,446,697,437,231,890,450,203,106,336,1000,168,182,39,427,29,706,143,485,979,583,133,140,618,49,130,324,942,310,487,774,568,846,789,962,809,901,248,762,599,538,373,972,426,925,211,212,866,213,243,999,156,193,185,500,394,290,881,972,876,679,313,806,199,145,702,336,667,701,295,601,388,446,789,825,656,383,690,442,536,99,7,804,237,116,999,787,556,454,820,786,948,658,986,871,760,405,746,857,694,944,339,317,892,399,347,2,791,759,783,501,266,998,684,869,649,707,821,632,212,463,837,642,277,30,307,15,57,953,265,784,957,250,106,760,606,942,788,809,859,304,888,476,26,409,405,752,583,170,632,453,152,602,283,474,775,520,978,516,710,353,340,474,53,915,840,340,642,445,675,344,160,402,112,307,67,638,929,950,826,42,44,650,383,162,354,651,948,965,940,901,96,33,13,263,126,961,730,334,797,918,52,433,734,256,722,93,826,495,473,85,941,439,843,83,476,355,413,376,980,813,641,321,755,246,377,915,724,750,81,296,10,211,15,978,369,372,174,43,120,591,154,155,266,2,202,695,878,186,222,712,90,542,204,843,695,370,804,55,154,248,869,689,823,63,796,338,374,832,94,449,384,556,598,961,463,279,756,691,499,291,89,374,558,82,292,24,356,408,65,274,378,137,22,601,59,791,417,249,547,238,341,738,697,236,460,941,681,951,654,631,802,477,172,293,532,797,688,104,349,819,212,399,660,476,336,800,630,175,311,109,562,85,863,510,110,389,484,704,210,1000,17,927,820,548,217,867,543,147,571,494,641,72,308,395,679,59,957,903,279,152,369,104,290,163,103,315,739,910,355,68,208,979,858,297,349,102,798,166,220,148,250,313,426,262,451,123,969,104,907,408,399,788,804,629,79,375,16,414,539,845,120,938,870,368,618,165,502,308,82,702,404,123,571,641,610,172,880,340,185,377,547,749,24,362,482,641,929,92,58,486,676,198,299,871,644,71,861,300,129,762,894,31,233,573,124,721,148,282,240,40,559,96,836,173,201,211,223,880,65,431,705,320,500,617,115,834,64,667,977,876,584,571,86,804,373,124,671,703,771,266,425,124,424,635,199,334,816,736,184,505,877,844,101,329,129,105,413,651,780,679,135,223,586,869,480,24,56,783,981,39,364,739,851,783,766,375,331,894,203,153,490,414,902,976,589,177,513,334,279,582,347,1,984,535,399,632,797,819,208,329,973,226,278,183,171,632,424,631,96,33,994,862,878,838,662,340,792,867,75,61,665,388,194,692,987,645,360,214,592,719,768,860,208,53,357,804,66,182,162,467,717,20,288,469,746,729,792,123,639,388,268,281,768,910,37,433,499,959,903,995,140,62,630,443,508,390,857,208,76,182,649,103,403,828,341,888,70,998,868,769,914,664,336,187,415,911,584,91,476,810,204,53,682,348,645,223,766,895,291,612,395,942,396,400,90,823,343,248,344,412,991,311,856,117,356,307,968,536,539,967,888,148,130,714,548,635,721,251,158,209,992,432,923,431,725,998,889,174,670,977,743,674,366,134,522,765,321,704,680,620,203,263,696,712,74,609,712,388,857,270,854,767,654,600,397,555,160,71,797,395,989,316,825,993,466,445,421,686,825,552,672,613,164,179,643,358,596,476,37,167,124,630,671,184,986,633,447,897,257,266,600,325,584,952,524,831,56,468,864,288,795,271,980,416,835,636,158,291,271,650,90,794,85,980,171,111,597,812,92,821,990,785,283,390,116,507,40,823,485,106,594,603,141,85,435,449,423,399,799,153,711,53,816,295,688,120,780,358,325,187,457,343,244,520,250,648,259,196,812,956,95,769,616,989,706,691,353,223,654,875,11,281,438,32,6,663,933,407,19,238,576,510,820,681,232,365,457,677,345,367,435,101,233,129,178,139,91,906,539,509,818,289,417,295,972,885,752,992,654,642,920,506,559,665,524,38,796,423,793,872,8,217,502,229,98,370,419,130,247,27,244,625,493,950,469,850,667,339,282,484,83,735,320,469,955,620,893,949,914,793,646,547,31,441,675,821,331,262,863,920,428,336,961,934,988,582,931,845,453,217,968,409,567,727,309,607,648,923,977,139,842,143,202,29,440,936,659,590,599,172,278,810,158,341,842,664,80,326,59,381,744,30,586,466,388,539,906,868,458,403,657,666,229,990,596,543,680,315,554,64,263,592,746,297,214,11,990,381,759,614,11,980,913,518,616,637,112,646,56,229,910,620,751,860,166,435,212,46,789,869,193,461,188,4,938,725,911,38,797,919,192,652,373,71,970,872,567,860,662,420,633,427,863,796,495,274,49,332,163,395,863,128,553,802,717,626,773,922,501,866,809,189,274,74,232,229,161,300,670,374,586,891,870,207,338,263,906,407,846,162,526,992,978,867,90,791,44,705,5,349,912,402,632,100,365,714,304,717,10,950,265,133,940,458,692,180,682,683,454,589,456,577,927,593,410,544,796,716,738,692,817,681,165,435,466,793,1000,596,546,645,730,584,849,638,555,253,579,965,368,925,860,840,194,344,211,402,502,228,850,650,349,769,433,392,789,8,497,432,620,757,77,861,296,171,961,832,538,264,46,416,958,956,31,953,596,545,292,427,947,93,767,543,339,531,306,80,722,504,842,795,736,119,398,495,294,771,138,138,75,515,234,615,967,226,247,626,696,99,719,56,541,775,383,962,223,75,258,396,713,708,725,988,983,63,303,318,846,372,124,718,120,468,38,235,150,763,606,974,445,503,849,898,221,455,798,48,229,228,853,522,673,283,434,158,888,336,925,33,423,577,268,689,96,489,118,62,43,951,33,908,224,678,549,368,386,439,797,501,616,938,348,364,403,294,25,85,436,98,752,3,626,584,838,950,895,757,978,622,617,825,573,138,788,63,2,401,343,766,709,668,982,911,375,456,605,747,39,723,229,365,450,518,906,126,73,336,529,87,567,895,56,215,895,364,313,119,636,88,46,786,156,576,92,238,215,717,204,258,829,998,160,566,63,179,244,996,331,880,477,83,608,604,911,931,472,261,555,3,473,408,617,347,942,65,861,108,904,349,346,92,188,881,840,28,658,70,10,451,967,920,683,502,848,437,446,231,263,273,690,820,651,734,102,908,312,824,837,814,809,691,473,972,541,335,672,375,996,986,427,462,153,842,256,117,720,393,446,369,100,695,450,784,367,2,946,122,856,718,726,118,118,192,426,59,905,488,159,896,714,125,99,474,464,230,605,124,9,278,60,259,482,651,812,887,663,990,426,358,924,864,182,920,473,642,733,843,125,1000,41,665,934,804,334,965,211,291,390,223,971,261,66,631,819,835,200,823,605,304,548,742,592,956,308,136,773,756,599,174,794,92,917,902,376,237,179,238,583,895,954,608,702,159,934,962,312,94,125,573,908,83,452,447,543,287,62,337,502,957,849,822,489,818,145,642,685,591,678,116,261,249,677,302,281,675,372,432,591,227,436,738,861,614,710,923,970,381,128,459,593,178,449,896,32,544,830,498,37,741,233,30,77,252,662,494,813,171,190,966,311,237,41,475,564,818,216,633,389,209,205,572,312,391,184,209,263,438,544,185,966,141,390,582,144,540,813,603,819,418,32,155,897,137,622,864,997,174,760,255,666,335,473,754,211,651,21,542,226,876,333,448,271,180,806,545,412,193,890,312,459,242,604,557,499,227,680,456,94,62,155,572,651,162,166,471,891,422,667,892,885,690,708,563,524,62,50,856,742,266,65,534,290,266,590,626,546,951,912,619,840,474,862,729,760,795,100,704,864,561,303,299,422,844,492,58,139,789,456,485,700,15,770,113,557,507,407,985,508,839,719,379,660,499,119,856,271,537,299,425,471,717,504,790,4,848,534,491,814,390,911,981,614,643,803,678,651,669,577,847,178,501,452,691,740,135,965,902,737,222,762,308,996,653,962,31,370,63,247,417,960,849,749,937,885,340,957,353,802,720,68,967,581,728,285,998,552,46,877,987,847,852,733,442,810,959,533,610,789,805,236,111,518,21,828,658,557,504,195,161,506,295,595,515,346,934,687,591,318,358,163,909,800,325,179,612,814,171,827,276,430,256,519,594,679,547,740,206,630,233,185,504,212,883,533,804,156,565,877,506,564,245,299,237,520,929,192,634,852,79,202,359,335,118,843,735,391,103,567,148,374,500,791,651,477,992,994,591,994,108,528,354,150,406,766,155,185,188,815,331,992,441,744,249,390,153,265,335,529,33,744,256,647,369,955,721,528,713,739,850,280,226,467,255,800,910,230,60,224,974,987,92,406,563,679,371,973,235,77,748,738,957,312,836,787,315,48,351,481,518,791,130,117,887,634,826,774,557,64,847,551,669,552,277,398,153,356,243,907,307,258,922,239,839,872,772,543,13,445,623,686,748,94,756,643,874,551,242,476,326,528,842,461,820,28,484,189,660,340,132,619,900,814,445,62,492,121,9,776,68,313,200,659,196,960,110,419,289,802,27,422,850,760,870,319,747,872,994,990,977,796,250,464,512,238,114,27,87,337,607,293,431,368,689,533,472,815,591,347,338,216,216,508,538,583,429,501,777,547,890,138,909,872,496,440,11,687,423,366,282,568,683,836,703,167,398,81,932,323,931,628,858,666,553,947,832,757,181,693,460,704,936,671,879,338,71,289,400,437,866,524,250,77,84,599,789,145,38,621,215,999,723,487,923,94,466,34,814,556,493,958,199,157,810,507,365,872,993,186,52,981,173,659,188,194,118,159,614,998,344,238,67,630,178,357,966,948,673,198,328,174,79,253,852,200,203,623,597,3,794,781,847,158,81,850,389,959,549,751,884,784,908,77,748,720,539,955,268,327,455,688,909,348,465,711,914,315,572,191,209,494,485,904,785,47,475,656,667,385,411,310,32,799,350,593,293,409,171,100,859,499,479,965,587,277,576,383,581,49,842,824,413,748,689,38,358,789,474,971,355,445,292,219,574,100,329,11,329,491,744,181,964,683,223,960,571,328,898,882,246,584,424,747,384,659,13,948,642,744,644,589,389,764,116,674,801,568,255,298,835,447,799,396,57,618,558,649,130,446,456,763,528,189,823,553,57,881,669,807,576,583,628,144,217,618,655,366,351,78,147,266,530,447,697,322,226,501,244,808,973,812,993,459,858,808,71,36,188,413,81,535,108,397,520,546,966,441,864,501,993,900,760,580,674,642,931,208,271,991,700,437,111,193,406,106,325,821,277,884,243,148,385,953,223,385,144,936,686,461,219,431,878,530,400,604,530,675,142,801,301,279,781,215,192,775,70,1,606,941,718,503,421,77,542,367,62,315,501,542,727,45,108,668,255,205,897,966,309,160,95,299,231,196,29,800,737,181,253,922,872,272,475,794,174,56,835,164,193,241,694,544,317,835,640,254,497,830,387,31,455,914,66,422,678,405,236,713,557,763,154,135,687,322,665,286,118,373,884,939,184,492,288,695,914,801,790,614,745,583,380,673,83,599,242,761,398,462,618,92,22,688,861,772,761,592,655,559,156,309,668,816,996,362,187,786,542,483,154,881,80,902,272,518,493,293,275,632,13,384,366,578,967,514,180,953,656,403,206,117,754,424,33,843,451,13,638,934,201,132,705,927,494,993,32,249,560,139,918,691,531,716,696,212,533,951,704,317,817,321,265,560,280,295,236,633,539,164,500,403,719,452,814,608,574,608,833,280,184,977,521,882,279,729,719,947,733,151,597,517,933,586,336,38,658,665,833,945,398,886,886,245,382,537,640,689,173,882,465,531,455,446,826,852,66,202,180,628,78,701,675,132,84,875,934,7,104,520,602,988,950,286,985,871,649,471,490,204,501,108,250,61,818,207,861,376,933,720,366,706,193,887,998,694,759,866,937,452,841,882,638,192,763,179,27,423,406,560,313,709,80,374,834,765,830,563,133,120,872,366,958,554,204,301,116,335,609,524,316,383,260,243,873,697,969,39,245,62,15,886,927,499,820,652,858,12,547,604,264,820,67,613,193,241,88,880,90,783,567,525,538,841,808,931,859,417,859,592,62,143,569,145,454,237,279,10,17,277,589,72,369,161,265,757,285,101,759,382,964,439,772,53,552,822,493,34,880,628,717,256,316,871,268,294,51,546,512,630,301,647,770,907,56,109,769,361,118,241,242,435,852,133,9,316,546,580,631,914,822,964,789,307,515,458,275,378,933,687,756,864,698,346,218,32,240,950,742,163,973,275,719,219,503,376,347,336,131,905,293,123,333,520,885,5,531,487,768,499,446,175,252,554,593,518,133,643,478,24,879,495,86,266,841,225,933,680,644,660,659,984,958,625,469,391,102,955,751,299,46,740,501,181,321,22,669,935,719,365,935,348,691,622,14,256,706,862,207,235,336,582,614,221,507,766,720,386,461,182,792,836,429,340,49,530,313,76,481,800,992,1,358,676,780,60,827,893,536,643,785,592,579,415,346,957,537,597,45,74,495,486,201,747,502,418,832,137,200,749,224,233,973,952,43,621,459,984,567,749,551,405,364,835,302,276,360,607,337,250,188,489,685,111,991,392,206,334,36,153,109,870,388,273,67,251,560,327,557,237,341,332,406,379,667,424,401,541,244,260,428,840,89,70,211,246,563,313,545,186,992,420,14,401,67,266,882,887,730,63,434,230,96,935,393,959,728,273,648,435,39,445,619,964,206,111,560,690,107,44,973,651,963,125,278,17,943,209,38,684,393,813,822,595,711,218,706,144,18,823,971,472,725,323,431,258,640,653,840,458,970,300,31,459,340,933,451,62,789,690,576,666,293,213,818,245,245,416,745,145,881,344,414,193,866,485,780,466,210,522,989,276,863,574,777,771,173,209,109,855,408,507,901,62,617,326,62,139,365,770,259,615,940,999,530,103,252,715,167,118,7,587,683,23,699,830,56,613,424,7,637,32,26,493,449,226,306,913,174,718,809,358,854,69,984,254,994,293,989,283,796,462,399,330,411,124,367,200,617,468,745,299,957,482,691,926,634,98,276,158,373,278,21,816,519,229,557,737,536,136,762,136,769,209,684,991,949,918,836,404,771,787,384,364,427,808,703,407,107,409,533,481,960,938,37,738,826,350,429,457,146,229,70,214,928,110,242,190,457,517,414,114,328,707,967,809,107,570,578,694,415,459,208,429,443,441,446,13,712,613,121,871,597,411,553,15,660,625,650,856,394,169,660,439,501,180,472,919,575,698,175,857,462,635,71,51,707,506,166,379,924,721,641,124,366,73,965,775,100,18,376,826,461,342,169,169,200,852,922,864,576,909,122,14,931,826,259,548,133,231,581,683,927,65,212,773,176,699,616,794,503,133,973,464,206,747,520,190,58,498,538,471,498,407,765,247,807,543,197,957,795,115,918,782,791,727,184,294,593,993,162,419,900,67,337,876,661,793,776,560,779,872,681,606,67,375,857,244,533,986,301,371,771,861,372,381,495,783,126,765,958,731,601,644,257,209,97,629,888,835,92,929,61,177,778,174,474,388,837,883,831,42,527,757,606,615,861,8,719,779,466,920,918,73,325,783,240,724,116,164,377,544,477,738,359,910,635,660,371,527,44,324,156,990,529,172,748,547,257,615,17,72,615,428,427,388,717,179,643,886,756,956,199,910,288,472,270,675,325,655,518,976,423,723,872,448,333,307,285,837,611,312,867,873,646,21,493,197,214,653,4,569,129,10,318,379,626,447,682,952,505,324,582,916,770,753,358,249,574,835,838,45,894,588,772,334,623,921,711,691,559,634,455,566,186,819,856,837,599,325,958,967,768,681,798,365,77,952,252,714,329,664,485,637,904,223,370,743,87,961,518,973,172,99,625,856,219,45,684,724,370,578,206,108,418,781,671,63,956,761,816,266,200,415,537,390,279,352,809,918,960,14,265,248,353,72,902,254,857,753,259,41,654,128,277,15,795,862,330,332,704,633,544,992,331,461,421,597,349,599,961,804,410,877,32,459,657,76,150,311,374,301,833,138,984,989,449,672,89,8,310,937,374,126,298,697,439,207,594,948,762,317,210,310,791,649,368,946,157,204,233,735,449,278,447,831,790,431,391,135,763,618,621,813,65,266,192,307,397,376,55,827,360,272,550,756,713,112,802,568,363,719,834,266,102,69,82,367,502,459,598,209,346,937,714,419,990,711,602,140,707,864,266,257,608,433,26,504,802,690,514,391,755,305,679,635,908,800,746,24,764,696,864,439,12,496,352,235,727,538,680,176,818,108,756,675,883,790,967,370,298,738,904,156,755,255,737,998,169,65,681,532,764,42,868,636,172,361,86,111,502,17,691,932,568,526,525,797,769,73,250,183,926,965,753,98,546,937,124,963,923,717,971,719,89,455,466,77,128,126,898,504,970,365,103,730,322,130,80,375,407,118,993,925,837,477,149,763,192,918,344,646,306,928,72,361,202,102,79,367,968,771,8,694,112,738,29,838,385,87,457,389,927,26,740,739,192,244,467,557,772,367,5,109,239,453,31,921,787,72,993,674,448,955,4,629,121,772,141,87,905,966,357,463,608,555,931,319,911,2,286,701,707,18,208,708,493,959,37,312,944,821,53,952,553,552,977,17,962,314,757,110,790,53,368,605,866,808,611,120,997,291,984,564,19,308,426,331,638,415,533,52,951,611,888,517,61,495,197,209,270,871,845,129,863,658,589,234,827,820,965,739,637,865,420,42,228,751,467,91,26,50,654,392,480,19,759,952,585,224,423,837,590,41,839,810,898,835,147,99,720,640,457,161,408,659,924,568,284,621,694,371,112,593,616,69,190,784,46,665,929,469,233,920,258,8,499,709,974,498,973,482,351,432,207,388,46,983,800,116,322,662,99,41,43,977,200,775,349,270,18,588,385,299,11,276,828,658,464,172,40,513,277,38,886,609,742,838,357,849,955,285,514,551,788,923,703,154,170,483,695,636,484,649,898,612,673,171,39,400,192,408,872,876,129,727,587,422,456,858,758,134,808,186,115,235,407,429,348,696,953,64,242,472,667,522,840,369,495,364,980,367,405,416,428,410,828,246,653,943,172,982,882,466,429,62,953,666,55,890,23,603,408,679,915,817,987,81,427,186,228,315,712,778,772,83,18,603,437,490,991,634,610,60,351,543,893,702,938,858,246,259,758,389,591,871,916,786,781,48,569,487,216,384,507,173,765,726,921,16,673,458,741,636,632,696,589,423,97,411,748,273,589,385,801,78,756,334,894,177,754,741,407,782,574,720,901,474,608,259,401,309,429,23,170,194,654,824,192,284,784,193,174,455,528,178,186,840,26,964,248,471,621,251,646,493,881,104,5,956,246,701,737,278,822,490,128,547,853,807,205,758,282,822,289,368,391,269,329,111,177,373,121,509,625,366,32,635,708,989,671,81,698,496,975,495,144,750,834,692,319,792,318,945,4,236,767,281,152,326,962,472,270,736,827,175,705,780,478,694,980,222,715,319,776,459,699,866,1,116,492,343,941,669,339,734,286,73,906,914,43,921,188,992,59,680,176,438,300,676,829,641,13,666,875,258,172,137,345,34,801,377,221,510,745,6,883,700,293,960,260,759,29,896,150,875,203,836,852,171,432,754,933,832,166,267,514,771,312,951,401,399,137,230,330,196,187,914,807,856,840,987,848,548,22,595,895,373,567,297,718,989,569,913,686,860,335,178,654,634,166,989,945,174,482,640,119,54,301,717,102,364,136,547,25,566,164,850,590,460,297,366,905,328,676,925,148,238,480,333,628,172,164,841,391,139,290,428,982,243,411,307,396,750,806,164,658,548,161,368,277,406,246,499,710,245,745,463,974,565,301,649,251,955,508,404,525,714,821,789,71,526,185,51,843,93,936,4,93,31,459,199,724,128,790,815,586,655,976,548,938,873,714,122,199,706,850,410,478,374,230,212,923,357,889,785,625,160,182,523,504,312,494,878,400,659,34,19,461,635,960,195,211,437,975,75,416,391,26,19,159,118,812,787,423,629,869,693,179,209,68,554,475,613,562,350,263,590,230,926,804,552,751,12,660,366,746,544,689,847,802,241,534,751,217,520,617,452,438,285,686,473,222,734,158,707,590,53,302,857,456,971,757,781,933,304,270,713,311,912,101,980,913,910,599,751,54,613,113,217,673,326,349,175,752,28,126,484,429,445,723,493,705,735,813,329,221,985,908,98,810,90,625,926,532,544,975,691,581,812,609,172,568,654,432,776,55,831,45,409,31,341,30,37,294,211,486,540,836,262,240,822,81,681,449,590,153,959,468,762,281,713,307,15,851,127,841,802,668,78,157,316,231,359,948,771,276,235,538,529,947,991,954,771,201,610,336,639,979,494,30,691,309,766,558,634,640,602,491,269,750,648,884,146,689,526,172,636,396,279,722,604,53,519,110,723,335,499,354,27,689,74,335,846,401,394,295,25,221,480,749,445,388,764,691,513,748,296,867,786,52,87,440,460,879,761,173,86,214,988,672,28,95,566,596,462,696,697,282,281,725,748,469,417,717,721,514,2,321,76,602,21,9,407,909,150,847,981,703,366,397,466,722,129,358,827,486,107,334,293,122,649,263,12,816,1000,916,213,632,796,711,227,116,268,969,471,656,911,694,709,577,163,248,960,848,857,341,1000,31,219,3,997,433,402,569,843,636,279,265,995,170,204,304,642,668,879,623,177,137,311,290,218,795,982,886,41,109,125,352,436,212,400,17,828,657,505,774,894,574,264,212,170,110,544,233,162,216,371,35,410,895,478,968,746,793,672,399,965,676,159,969,142,939,657,994,37,578,268,9,336,647,481,325,748,897,748,525,425,311,855,797,233,923,460,317,854,863,984,578,396,1000,708,745,799,787,205,1000,873,669,377,360,493,736,94,201,694,649,600,640,103,280,841,434,201,364,885,213,376,12,787,447,460,967,693,298,894,911,144,653,324,594,762,652,682,456,908,677,733,621,362,601,477,496,560,848,776,19,498,824,316,144,492,600,839,365,89,747,405,544,808,90,740,196,62,108,964,916,399,208,621,621,209,976,587,400,273,622,273,29,74,439,945,917,691,479,268,272,982,341,987,469,834,688,849,467,918,550,608,44,878,650,520,18,684,810,167,965,602,449,962,837,244,485,93,54,159,700,289,650,822,198,325,921,188,66,716,506,185,640,460,164,309,718,873,226,802,148,332,927,953,863,709,962,18,379,231,791,635,452,260,372,385,391,485,213,153,656,452,127,674,128,551,21,888,73,448,222,662,774,173,622,427,403,978,790,372,941,398,681,589,220,480,285,323,593,57,977,876,450,877,839,643,618,15,550,839,973,121,153,733,322,912,804,346,502,188,291,654,747,837,471,55,12,109,951,665,122,771,210,52,183,980,680,549,43,442,785,952,360,353,267,470,449,525,981,303,961,734,31,967,774,164,690,122,711,418,138,240,690,834,593,257,114,661,704,928,380,259,308,251,337,628,86,171,129,341,715,905,444,969,73,271,348,80,577,516,972,540,643,791,862,132,342,456,117,402,566,734,539,668,539,183,949,924,481,384,772,817,36,718,398,525,550,59,590,280,462,675,245,431,846,787,877,922,547,937,159,544,524,184,989,6,697,45,375,180,601,752,646,792,627,784,50,331,906,670,875,358,381,303,171,961,531,322,617,537,303,917,334,133,949,497,443,389,261,139,340,444,695,459,505,47,852,465,668,251,524,872,323,926,892,552,871,470,310,360,217,159,51,356,396,360,271,581,861,188,406,58,822,896,835,73,18,557,351,817,408,900,332,60,187,361,376,70,589,524,454,342,958,45,798,656,141,268,168,592,126,459,1,855,651,928,440,236,653,24,943,121,645,449,943,244,533,35,797,894,679,87,353,542,514,533,407,367,115,254,992,771,699,559,747,634,830,93,107,794,854,909,268,253,426,184,215,707,46,405,445,992,113,159,189,266,71,789,779,191,937,914,638,601,696,739,970,40,289,471,833,274,897,185,995,932,112,153,753,114,248,286,514,937,951,937,585,687,559,88,524,755,955,840,572,809,212,267,68,945,745,535,286,105,69,54,727,553,656,345,556,971,555,304,442,551,948,170,846,638,20,837,519,698,133,206,652,768,577,777,34,597,868,219,659,44,692,1000,172,461,469,166,141,820,633,521,435,172,575,820,8,368,609,776,640,220,583,370,518,419,736,623,188,244,311,197,957,764,785,327,238,616,656,962,155,951,309,4,204,533,173,382,801,613,514,153,5,603,902,514,202,615,856,920,596,931,867,328,91,491,777,287,438,57,171,653,648,885,559,924,225,429,325,149,581,955,291,343,313,230,648,567,917,67,415,48,508,104,666,745,180,137,1000,220,922,906,225,664,981,1000,635,247,1000,690,961,719,82,542,339,242,843,516,713,315,521,980,70,571,540,359,153,194,222,902,448,571,546,139,137,365,659,351,477,852,774,474,220,586,827,976,122,968,31,926,481,939,440,792,858,631,119,905,755,59,525,917,206,602,422,277,29,351,816,741,195,833,759,743,934,115,599,671,318,266,823,538,199,754,249,806,829,90,535,305,323,356,947,252,195,24,48,224,442,586,421,327,180,292,343,727,679,578,280,358,389,206,590,707,258,635,280,186,612,973,25,541,81,254,458,358,914,518,500,36,242,945,94,618,763,957,93,869,858,415,232,668,154,36,713,136,33,665,241,753,991,174,851,447,64,834,213,320,561,239,50,525,472,433,936,173,18,752,550,697,278,428,175,43,887,641,733,520,826,660,874,997,49,464,658,659,307,255,561,60,125,320,731,775,998,477,178,562,697,104,986,425,863,32,502,745,26,228,415,408,178,365,44,749,44,543,530,252,262,530,456,523,889,529,691,63,196,569,789,9,478,624,839,391,675,937,1,109,169,81,262,70,737,904,847,456,417,395,594,38,628,855,416,209,712,770,929,555,235,91,761,317,15,465,477,625,713,568,486,156,323,618,853,169,272,197,943,279,221,129,554,648,118,717,784,854,453,58,335,161,262,80,392,383,596,357,681,705,950,618,935,386,369,269,761,219,761,970,66,4,539,868,108,42,5,927,8,877,915,386,831,214,830,238,370,367,153,833,394,274,162,565,521,223,529,482,240,339,558,389,616,72,650,600,732,420,728,357,431,49,597,548,612,260,185,891,13,702,912,491,939,166,305,351,25,524,779,338,216,800,966,586,472,260,343,68,463,927,949,214,183,928,817,836,957,613,800,748,795,790,230,731,198,103,44,290,139,623,125,303,850,611,840,126,323,412,398,9,352,630,446,186,355,538,958,947,212,952,384,793,33,519,551,325,431,915,819,132,884,771,817,639,20,851,773,696,670,480,245,254,291,597,349,927,934,735,806,178,376,203,697,573,847,712,811,113,22,305,799,991,198,75,320,499,812,608,156,375,372,499,137,863,857,779,401,894,91,37,859,305,887,707,349,477,734,516,760,805,349,723,781,253,164,701,947,340,474,636,492,516,893,16,50,580,345,613,47,843,967,89,691,123,346,844,862,778,243,917,576,169,55,389,288,78,593,800,814,862,17,331,851,306,870,806,487,264,69,976,717,643,129,953,145,979,392,177,356,547,620,43,379,268,969,927,392,281,304,239,547,948,519,848,49,44,287,953,784,44,112,899,548,659,850,67,795,402,837,25,39,691,517,182,767,320,616,484,515,638,21,700,429,645,683,300,408,250,221,846,850,793,130,787,965,633,927,62,32,124,753,782,669,958,497,454,406,879,36,829,525,609,317,718,13,690,1,239,887,188,221,815,316,551,639,178,744,905,43,137,985,24,700,742,733,464,642,840,401,407,579,382,254,485,250,79,459,849,669,803,688,315,782,556,997,793,829,123,160,599,437,536,289,557,788,768,541,155,374,675,376,820,245,431,360,225,756,309,572,101,863,459,525,408,606,671,275,274,57,272,508,242,137,156,44,714,239,811,526,497,916,607,4,393,301,764,969,939,165,131,217,880,192,667,870,865,926,810,945,662,453,367,303,130,435,150,991,720,194,775,251,958,966,856,443,916,851,609,432,442,129,972,848,838,194,303,360,400,834,2,298,227,672,930,100,27,189,765,468,49,953,988,577,478,536,803,639,569,148,964,400,109,273,827,333,197,902,892,163,584,616,632,431,700,703,847,257,671,116,895,35,836,732,137,304,388,853,133,67,308,529,816,18,719,318,679,425,676,289,620,963,740,341,520,604,595,645,423,239,19,483,832,532,682,194,753,414,553,871,744,929,334,13,954,563,542,611,116,846,601,808,751,804,921,300,621,567,293,433,425,138,57,235,323,538,818,526,821,181,490,827,678,335,935,223,392,696,290,294,600,505,628,307,742,903,511,173,298,801,965,533,856,369,495,949,753,410,881,487,352,977,29,313,62,520,452,725,49,142,565,414,658,101,844,80,155,31,793,944,48,783,193,330,427,190,17,537,187,237,241,124,648,113,879,235,706,799,115,273,578,398,562,677,900,278,536,613,106,617,433,998,531,657,812,12,651,33,274,547,506,260,538,781,403,222,305,79,391,790,848,722,131,200,183,889,466,201,301,319,740,392,780,177,685,931,879,108,337,495,133,202,88,12,20,612,169,324,176,771,601,376,417,638,672,668,136,901,718,496,608,411,848,822,747,152,772,516,766,810,225,826,380,369,732,138,444,576,857]`

	t4 := make([]int, 0)
	json.Unmarshal([]byte(testCase), &t4)
	fmt.Println(numsGame(t4))

}

func numsGameBad(nums []int) []int {
	output := []int{}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			output = append(output, 0)
			continue
		}

		subArr := nums[0 : i+1]

		output = append(output, getMinCnt(subArr))
	}
	return output
}

func getMinCnt(arr []int) int {
	min := math.MaxInt
	for absoluteIdx, absoluteItem := range arr {
		cnt := 0
		for i, currentItem := range arr {
			distance := i - absoluteIdx
			cnt += absDiffInt(absoluteItem+distance-currentItem, 0)
		}
		if cnt < min {
			min = cnt
		}
	}
	return mod(min)
}

// 好的

func numsGame(nums []int) []int {
	output := []int{}
	// 预处理，散点叠加 f(x) = -x
	preNums := []int{}
	for i, n := range nums {
		preNums = append(preNums, n-i)
	}

	// 遍历求每个切片的移动次数
	for i := 0; i < len(preNums); i++ {
		if i == 0 {
			output = append(output, 0)
			continue
		}

		subArr := preNums[0 : i+1]

		// 求中位数
		midLeft, midRight := getMid(subArr)
		cnt := 0
		// 长度若是奇数：中间是中位数。若是偶数，左右需要对比
		if len(subArr)%2 == 0 {
			cnt = getCnt(subArr, midLeft)
			cntByRightMid := getCnt(subArr, midRight)
			if cnt < cntByRightMid {
				cnt = cntByRightMid
			}
		} else {
			cnt = getCnt(subArr, midLeft)
		}
		output = append(output, cnt)
	}
	return output
}

func getMid(arr []int) (leftMid, rightMid int) {
	sort.Ints(arr)
	if len(arr)%2 == 0 {
		leftMid = arr[len(arr)/2-1]
		rightMid = arr[len(arr)/2]
	} else {
		leftMid = arr[(len(arr)-1)/2]
		rightMid = 0
	}
	return
}
func getCnt(arr []int, mid int) int {
	sum := 0
	for _, value := range arr {
		sum += absDiffInt(value-mid, 0)
	}
	return mod(sum)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func mod(i int) int {
	return i % 1000000007
}
