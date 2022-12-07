package main

const Input = "WjmsdnddnmQPZPPJPL\nbQllTtpBlgwtrbbCwfZcfSFPSfLCSF\nGgVgQrlpphBGrlVGgTtsRHRWVRMzRdVsqdnDnV\nMMTcbpnfNGQbMjgsRwSzRptRzz\nlPqCCqQdQqQmCPRzRVSwtzgqqwqR\nlrDdllPdBWdDFQFnMbDNDn\nFldWTldlpBSLzvpnpSTpWbDhbHNDPHhJcNHNDwbH\nqVCGQRGrrgMQrJhPNchgvgJhNc\nfrjGfMrGQMjsRrRQjvQGmrQQszTTpSLBznlzBlLLSBLLSZTn\nTPZZZMTTbNTZNtTlTbjPVRVGzpGQLLzdGgmslhzSzgLzQh\nwrfwDDvcnFvCfrrSQsmzGGQQdndsGg\nqfHwHCqCqCrFJrcBHCCJRmWTRTmVMttjRjRHZMNV\nvpbqnzbPmWLFjFLBnjZg\nNltQcCClQlcGQGtMTCRdGTGBFZRLZjFFZZhBPLrFHZFjHf\nlMtlGMwMsCCNlTwtsCPzzmPmmVpmpWqVWsWz\nVmWVSchSrScGtwlVtBnwBVFF\nRZZPRNpPCLZvZPZNCLbQPZNBFtttmwBMTMTtttCwtgwBgl\nRNNmPvbQQjPRQQNNHpNbhsfzWJqqSJcsGHhHcfhq\ndtJvcpccWvLDztRCRRCrCC\nqHVslPzPqHqzmPhTzmDFggjrHrSCNFFFSjgR\nqPVPMhszZPVhwBZcdpvZJncbcJ\nHBNLlBDtvLDHhHLvfwlFjqfQTFqqWfST\nddsGcggJncZVRdGCdZdcWWzSFrjzQFrfTzqfCjFw\nVdgscZmbZNBvbDHTbL\nGtSZQqpHpHfGHzzqHzHfSbPbnJCrRCnJChmjnJnbSh\nDvcTNTDWlNDWdlbnnbjmRhRrCCRd\ncclNlwBRvvTNccRlBNRspVLLZzVzVZQqfqqLpzpw\ndBSfHdZvMQMdNVpWRmWmLCmmtB\nrbTTrrjDcqcrrqrjjJGGclltPCWDmvtpmpPpvRDsCspsLC\nvlhjThlqcjrnTvThndddHZngMnfwNQFNFM\ntHqfszrgLsvgqtHrHtwVCGBRRjGSCwsCsmmV\nMdFMclPmcDQFDlDdlZPmbVpCbFpCbBSCCRGCFpbV\nmcDJdDQMcDTZQhhNNPldhhDtvHLgqtzgLvzTzWHWnWfvvz\nsVdGlTMMVTGCdsTMHHWWnNBzNWpNWCpW\nmPmjFhlwmwmWrpSrSWHB\njhPhjwgthtFFjRwjZgjGdJcdflfqMsZLLsdZdc\nCBRsTsBBzLCfLqtqBRPNDQglSttlcgDlgGGt\nrJdbrJjrdbVJZdVZCGrNNcgDQlcQrggl\nFpFbbppjFCdwmTvsvfzmTnTBfq\ncqhcWqqCNjGWqcqhGGZzngftmptLZLGZTn\nbrJHBbPVHPvSsdHrzQLzZgpfVzpfQtnL\ndHFPBPrBJFRzcWNFhWwN\nQTBTfQTZsjWDJBJd\nFgFWNqWGDPqlPllp\nCCvHzSWFrrtvNvNNHLGQQfbVRRfHZcZcVTTTnZ\nZCCHHCVRZzBZQThM\nnljDtcqnhcfbwjwltfLQMLLQQppJBMLQJL\nqDsqschsqblDqjcqtRmNVVdNsddNNPmFgV\nQQRnqGBSpQnMmSGmRQQFtdcbbtHHccjpTFcTfF\nwCNPNwNNWNgZHJHJFffqCjbj\nZNPzNWgNrsmzMVqsqs\npVWlMBWjlWWqspWDjdjMpMDCPtmmdbhtQtQtbGPCzChchz\nHZNgrHSvHwnFZnvgNvnwLPBGQHmCmtPmtBCGGCmQhG\nFrNfvSNSZLZJrJsVlRWWWqRBVTff\nPQctSHQDPSQcbShpFzbmFddpmdmR\nwqWVwvwNCJRhmdhwJw\nnMnWggVqRVZqHBSBsQsQGDSZ\nqppwrgZSLsVbbfvZ\nhhBHPQQChCDcPcsvvhbGbsllJTfv\nMDWmWFCPFWtgpRWjws\nbmRjdmrJRjhJdJLZBjTFfHGTtQFTSQBS\nwCNVnsspwsnvNDwnsDwSBHtDHHQGTFQFtMrQMS\nnWsNvqVgVcqdJrchrz\nMZlfqlmblmMRWhWNsjSQfh\nCznczgtDFnVtFBNSRNttvhQsNh\nSGDFzVrzVPrGHVnzCPVnlZlPwTbpTqlbqLMpbLlq\nlbbbGDlwLDLjvDvm\nFQfQnLTWVcPChtjmjWSj\ncfgzzgfgfVdzTdfNwBpbLbwdRbpZGrdw\nmwnWtbmdWdccwtgTmwnQfPqsqLQQJQQLsfQQ\nSjrZbhvBZzPHQqfGJfjG\nFlrMBSvMZZFBZhShMMdctWDtmFNDbDTmtmpt\nndHWprpqFTnnpdNFlhljzlGTwDGzlhGz\nfvZmmVVfJctMZsgMgmcBmsZhwQlGDCzPPZWDGDjCllhC\nWsmvmsmRNnpRHdbS\nLgZSvhvcsWtcWnjrFrWjjjnPWW\nDJDfNlDNRFlpMlFLFP\nTmqTqDHdmfwJRwfdQJLHzzSggHZZcGzhtZvcZhzG\nlbTpqhhgSlgtlTqSDzzLPPPrLGTTGnLm\nVwGfQGBGZWNnmDLPznLB\nvQdfZFWvFMMFwWGGhpGqJbpFtgbS\nhSvCvFRDwmzCCHrszb\njdMgfTblgjJTLLLzqqGscmPmlqPzHc\nbjQLTTbZMfJffZBBWdjBBQwVDntFpRvvDVvSnhppFv\nFFnFQndPqzmjHscmJFwc\nZrZrRgDphGGDZgRRBSBNMHHvmwJJSNJHcv\nDGrbbthfWWWDgtfffttfpGPVqnmblqPqCdQTlTPCzTll\nsHTsGrHpsftmddRRZfRv\nMqFcqcMQbMcVPCdFZCzZsRsFzL\nqnVlBbcJJbbMcbgQJMMbQlDWWwHWSWwpTsGHGrpDWG\nqpmvVVcGvVVcVmDsCfqTHLLJzTjfZzLZ\nrSSgwwnSRRBrQrZzZHCQrZ\nBFgFRgdNnSBbbnhSMBSNFdsWVtpWmzsGtvpVPvGWVDDN\nfBBRfJBzzMRGRzCBgWtbCWtbgHHHWdqt\nmvNcDcsDLLnnqfbv\nhpTsTsTrDrfpMSFFZjPFMrSP\nCNQGGDMFWGnWWvvNMQFPvrgzBLVBLszwgggLgw\nmppZVZtZRTbTTpRccVbgrSLzSPSzBrfbPBbSPw\nVtZJhRJcTpJTHlhtHZCFFWFjNjGGnFhCqDqG\ncppcZGcGgGpdTgSSnmpFMFrFzmwqwmrHwz\nRStQJNCNvfBQNjrqmrjFrBMzFB\nLJDNNJCDLcdLnSVd\ngNrBNSrNNtSjBndzmzlVnm\nMbfqfpCLpCsLqsLFSbQLfnMdcdRRcdzldlnljzncnJ\nQPQLqZqhSqhvtvTTWrNg\nDtrrcGvtLnrrvLrfctfHztrfQpbwwphpdSbbPPPwFSFFRPwH\ngggBqTNdTNjqWBlNmCqCmNggQPhpPPPbhhQpSBpJwbSJSQFR\nZglTmgWNgVZMZdsGMcvDcMdMnz\nbjtTFsPmmtpvVlQHlQJQnJjn\nWzDzwLCSLrrDNLdrSZRCwNzrlQlJMnJQJJVhbcMhgllwnJQh\nfRCZzrWRzzGbGvTmBPTf\nlVlfJVblPQbllflfLdJdvGpjnFRFqJFnDqpJjnpF\nHcwZMgmwWCHHCSwcWCcgSCtCqqpFppGDqvDnRhgnnqFDpjFT\ncwCHtGrCssWHCCWZZMbPPNBVbNfVbPllVszf\nDSpSnRwrZDPWsJdZ\nzjjlQVjlNZmCVCfhCfgFFfFFFqWJbgbFWHJH\nlNQMlGjQBZjCmhNMCChGzlVNrpTnccLLwcRwTSppSpprLRcB\nvLfvcgglbfLfgqdgNpPtzqDmPzmJTTztPCHT\nZWSQVGwQcWjSshGwVcnSzDJZtPPTzmzzJHCTzDtJ\nSjVSrWVhQVQhwrLMcrFbfplcflvv\nNgtfSRPnnRrSlgsPhnShDWQlMWpVBMMMpCWVBpCQBB\nLLJnvwJvZHZbHTbVCQBppCFJCWWQBz\nwHTGZmZvdvLvjLwZdqngNgsGtrDPSDhtNfDf\nDwrDRlrwmbSbRgwsSbRwGJvQGqjJqGNTJTNGTSGn\nPZdMZzCQFBZWWFQvJZvcNcjqNjjZJG\nWHWCFHBBdzzMWhPFtFdMzlRsVbVmDrRVVrtQplwVwD\nbpWbJMWpJbprfNMrBfJfprWhPnGtnHnLHjPPjLvsWnHGvGvj\nqZdgVVgDQhQZlwcqgDcchldjjmLtntmntPQsmLnLPjssnL\nRVRlgcSSdglZczdqbTCrhBrBpNBSBSbF\nSgbGvfbnGgmnNnnzqMqqHHRzbZBzZR\nTWlssdFwWdtswWPtTtWltwdVHlZZzRHZBZRzrprqHMpqgZrq\ndssdCWTFwSvgmDjDCG\npqsDnNzzZsdZSnDSpwjBCBWvgjvWjNFWQgWC\nlVGtRtLMGGfbTGTtTbQCQnBQBnBFgFQcgjfc\nGPbTbPtnPttmLTGRRbtmdwSDwpwwhZqmdpzDhDJd\nMdccRQMJvHdgZggvhjjMgHcHlWWqFFWmGqFbJWzzFLWlLmPm\nTDpSsTrtblSzQlGQ\ntfVNrwTwtswTssMRjMMQQddNMRCd\nGqGqGpFqqgDGFRqDwwqqmzpGTLPvVWMPVCPLJLRJJMLTlCWV\nrHrSbrsbQcbtdNHHHfdPlvWBLLWlvMTVlVLCJf\nrccbtthHSHsNHrrcttwqpnDFnMmpnzFnFhnF\nvQQQbRvlLjNNLLBzNllNHNBqGqhMWhGGhTqmPmqhWTFhRm\ntnsZwgSnCDrZSCDsfTMSGpPWmPSWmGmbFq\nCrVnfnCCtrCgfrffcrstDnJNlvJdNvzdBcHdLBJHvbBv\nnmQsMqTnLlmmpQZmTZcdHwCFSpHJSSWHSJVSWSHH\nvRgRRtfPvDjzDgDbsjzRvjfNNWWCSJFwrHCFbrCJWJHCSC\nGhsRBztDBgzRPstgzBLZcqmlcLMMlmdLQBZT\nCWfvvhfWrlllSSRrdQrQDQGQdTRr\njsNctMZLmMZLMmmmbbNZswZNqBTHPHzBMHHTMGqBRMRPDQqP\njcwjntLngngplgFhgRvJVp\nvchzqzwlhzRqzVZQwqtVPZLnLLbDnDFnbGLnDbPLDGWD\npBTHpdpHsrNBBsgdrdCpCpCgRFCMFDLFWWnFWLRRGWbDFGSF\nprsBHfggjpjjcRQlqvtw\nVDwzLQrDDWrrwWbJrVJwVrVQfMfSCNPMfSlMlPcMmThChf\ntsjFdsRsgtRmGZHpHRgBClSlGSClcPhGCfPlllll\nqBBFBZZpmgdFHdstjFJbzVwJVqJWWbrvWJDL\nmgjZmrqmdsmGtDplglJgRVVc\nnPhnLvnHLtLnWzzcNwwcchVJRflhpc\nSnnLntLWZsMqZrSZ\njcrNfnrNLNNqFgbDfCSgSQbS\nzPPHtMrGGptvTWPVvzvHRgQsbDsSRRCCQbtJsJDS\nwHzrWVzPwThGGwMHzTGGPGwhdlZnBndZLljNjBcLdZdNBN\nqNPhNqddBNhqvPhFvllNgNBHCrrCQnjpCfPVJnnJQJCjJj\nZZbZTZcmGWWMDWRSDnBVQCjVDffJjnnVCJ\nSGBRTTZGGcLSSWTScsmcMbGlhgwFqslhzqggghwhNvwwvw\nGCCPwpsBqNSsBPpSCrSshzQzLhTvQhqTnhtTVQcT\nJDjFJfMJgWbWWlDJcnvvhvtdLnjnzhjz\nFlDflbZfZgJgMgbmgZJfSpGCvvrGrRCpCBrZRprw\nHwqhgFGSMgPPCGQQQnvvcpjn\nBllbdfRBsBmsmZlBTmQCjTnNWNCmNmvc\nlJDlBflDdbbRlLbfsbZBJtbRqrVFPnwwwhVMHwVrgJwwrSwH\nLTvLtTFLCddFTTthsbVVmHHcqVHmWRcmHL\nlBgwwNggwMwNVbjBCQcCqCRB\nnCMCwZGGNGJnGhtrzsdDDndtsF\nZlZdJJplLZBDpJjNJlGjQCLmCQmTwVVCbQQbqWCT\nSfgFzftrnRzMnVbPPPQmPq\nFfmrRvgShchvFfghzRvgtvFBpJDpcGNGdBHJlpDDDJlDJZ\nLdNrLzjdWQnrDHsD\ntBZmBZtVZpldVMPRnsRQsnnsHVbRHs\nMlfldwMBBFMZTSFTSqLqcvSJ\nnJqBlvvBjHhBcqqRrGPrTrGpBCGzTG\nMfVCVMLZVZtQCdtLMtQtQSSSTSzFRSRPFpFRGpdTFd\nCCQQtQVgfbbQNggsNfQZZbHbcWmnhlhnvlnlHjJWvnhv\nSGmmGwVwnmhbhnhwhhwbdMgNNgjmvMDrJTCgmBTBBj\nztFWcWQQfcRzzRFllvjDjggjDDfBgJMBgC\nQtzcJtFtqcFQRRWRRQWFzFGpVnVSGLLGZVGVqddbLqGL\ntBmdmQtjMMqDLqBtttQMjDdwwgwccMMbffllgzncwfFflF\nTPVHTVsRsJVHVrVvHvrRhVJbfCgFbzwbCGgFlwgcCwbn\nhssVWnRrTVZSSZZqjLddWtLBddqtQL\nWhhtGZtZGQZmvCfCwtvhqgbfdDJdfjlSDlSlBBJDSg\nrHnpFHnDrdBggJngnL\nHVcTzFPNzTpPPVrzmtwtvvvmNQwDQvWw\ndStBwStGGBrNnBdrSSMzvjhgFcvvDcnnvPDn\nRLZbCWWJbHRsTHspZWLcDWcPPhzczfjgPjjvjz\nZmLHJqJsLJJHLRsmmmGSztQdQzmNBrBN\nJrmRVdvcmvvmvvRTdBVVfjFQLwjqLFLWFMqwcFjz\nDDhhttDHHHbHSnsDbHqMqzwQLLFwLsjjLLQf\nbSbPthtgGQPNPHnSDChCRJZZJRCZdCrCBZZZZvJd\nrqvVqNJpVVNwnqqTwthMMq\njsFRFDQRLQDQmsPRmQsmcQFMzGhwBGBhTzTTzHnRhnhBhRwz\ndCccQcQsFCmCQfbJbvZMNZfJrJll\nTGjrrTRLHvrQrFDCrmzzVm\nNwWqqhndWtzDQhCzVCsh\nCSSwNNwqgBBBBbGGvLTTvb\nfRBRBHCVRRzcCdZHvRvZVCZLNjtwtNwNTtLjNtTpTNttfS\nDZshMssZmTMjwjSLtw\nQPrJDDJsPDFmFrFDscHHRzrcvccCZRVzrR\nzgqzLLvlvdgpgrWpWW\nRnJmNRncnScFmZSScrJQQdbpGdHbWHtPHpBHsdFdsW\nJDfRcrSnmDSJcNfrNZNjvlwjhllhMzlDqMqMvh\nfDLzSMLhhtDWMvtjCRRZjCHHJjChHN\npmTNpVwPNbPwPBFRqRJqjCnFjZFV\ndTwpddsTbgbQBssprsgtvgcNLzWMctfSgcfWLt\nfbBsBTsNDhGBGZcLLLJJQffQLQ\nMpsCCMHClsHQqZcQWLqR\nFtjdCFzVljFlslVCpFrFjPhggBGDgNSTTgbGNmbGTr\nHqTfmsCFmPlGHddNVGpLhz\nJjjcQQJgjZvZZzzwgpNVGwLGgV\nnQnSbDDRbSQJQQpRZtZcZZsPPrFfWCPFWlrFsFPqmqFt\nTgTDDrCmqJDGLrhqLmLGqDQRFtttjMbQZJjtdtjFdsdF\nWHffcHWnlvvcSSWzPVvHpWWVRbFdQQtQnjwZFMwZtwddsZbd\ncBpPplVVPfvVGBDLGCBqmmLM\ndlMMmnmjvCCjJrrvMdgHcbcFbqFbzQrFbGzb\ntPhRBRZPtZRshTzRsNShRZNGDfGgFFbFQqbGHDHbqfGD\nTSPhBVsTwRBTpVtRZpVhZLJjWzLjJlJlJvmwWWzwWn\nWcvLLgLcczLTDtccbLcgzMMfPsGwRPjwfMwHMfMvMp\nQVmlPQCdJlJJJlFJJJnPQQhlwwMMMnMwNMpwMGwwfswwGMsR\nlrSCZZVFhPSZgzgWttWBcc\nbwbbZLlbwlJhBzFCgtTGRGQldQRmQW\nSSnpHnPHqpmggCWgdT\nPPscHHTfcsPSDVfVssjvwFJLBJFjFJJZFJLNLwrL\nnLgDSHgwRgGnHjjNfTRhjPVpWV\nBstQsvhQZQQbMvCvMPVNWpPcTjfmPmmW\nbrtCrtvtzrhdSDJDwh\ndTQTwgmZQbDzzMQCCl\nWLLtntFnfnRHbttnSRRzSMVGDDMGzVlV\nJsPhFtfbLWnsLPLqgdJcjmcwTwjcdw\nwMwMbMRRBBMLPBlhLRQlhPcWzgJNvJtzWNtJptpgjJgpBj\nTnmGGmVnFFNSZsnZqFsWzWjrTJzvzWvWgWtTWz\nqnHGnVSsqZCddnGCGCSNdDbLQPcLLQlDhPRMhb\nNGsBTBlqsvfQBQqsTLTFltRMmRwmmHmFtPSRhM\ngjZWJWCZdDpjggDdgnpWdZZJtFwFRFFRMwbbmRPtShnwRbtt\nzZpSDDgpzcDddjVWggJsQGrfQvrQcTGTGrTrqr\nhpJchhFWMpRDWHWcDGnCGrnGnwPTwpQnCt\nbmgddgmlmjjbfddgmmmNvGQLrtfTwrTtLtGTLQVQQP\nqqbmdZgzZbvgPDccHcZPhWWH\nVVrdQZZrZSZFgQTTTzggrVZVMlfBBfvcMBCBslMhBvsMBSff\nbQwnHwbNpwcsCwjMBw\npHqqPnJHqpPNJFzqTzQWWDQQZq\nHgwTDfgBwBgcRHqRRjHqHTzQQClSzvlzPVSQLvbbPC\nMnhtNZNnJpWpGhMQbCCPVSPLNmSPQQ\nrrhJMFJJZFJpHcjRLFRfHjgj\nvGvGMBlttBltvjdgbPsrsDWdjPPP\nHnJQHVqNmQHmZsDZPPrDWpgFps\nJSqJQVVDqqVfJNfRffGSGMBwGTCCCTlBMSBl\nPsFZPfGbDNbtQmCCmCBBbmmL\ndcRdhSrCqjThTRcTpLzHQzTmpmlgBz\nhhhcqCvwhhVhfPNvsMstZtsZ\nVvGwBBwvZtGgfZCqShnFFjSstCMC\nNlTRdvpDdTRNzdTHHnMssCnCnCqjSz\nLvWvPcWLpGwBwVVgVc\nbVVmSrLmLSJzTZMSFTBdMj\nQnvqRGGDvWpQWGDpvsRZLBzTRjBTtjRTPtBF\nQWGvDpqcvpGWQpGngqGQGwpLhJVNJcrbfrVbfbhHrmlVJHVV\nNGRGPZWZpblGcJtfssSSsbffCs\ngwRhjvrgjmwgnzvJJJtVCtHJqs\nrrmLrhwFFmmTMgFRjNZWNpZlZLppQNcDWP\nqsHZsHZrTBtZrHBNFCJGWrMcpcddWGJWLG\nmRDDzbPVDVlVDgbgRRvmwCcWdGvJwGddpvLm\nPnfDbPbVzDbVfjnnlbzhVFsspBSfFsssHQNTBpFqNQ\ngCmtbDqPVVVqggCGqTJjHMpMQfgMpMHQQpcM\nzNZnsSLzZzrlRhTTJMHppjfHSpvp\nnhdrBsLlRRrdTlsTVmCmGqDVFdtDPGwV\nFnqNfdGfgzmPLGmj\nblvVvbsRwgnzjCPcbT\nRWWvtvphnZQZwMBNdHfNDBZZdq\nDMRhDhdvnjhnPnvPMfdZSGTccGJFjGFFpFpFTbTpTW\nNVgVmtzVlLBmgztsBNmtgCmqFpJJFGGpGbrcGGsrcpbWPr\nzwBztLBzllQPDZvfQZfRfHSR\nmFCgPzmqgtPPqMmFWzbMttcRGvRclvHhWGGcZvclRfHh\npsSNnhnLGnwZHZGv\nhBpNLTNLBhsPmbCgBtPDbM\nJsbLLWLJRfQFnccmQhtvvPCP\ndwgrVwGpgVhCrrhPDHtC\ngpGSjpVdVpVppjjVZBwdCbFRWLzLMSRMbNzfzbWR\nDmMQMJmnmGwzGwwG\nZcLcgLgcRsZSctHFWLGfjjBwvbvBsvjpfhGb\nGgFPqFtLRHLFSHLRRFRHHtMnJVNCQCrJCJCnMJDdnqDV\ntqdqFqdsRdVdtHMNdRZHTZLrHJgrlZQPJLgr\nGGwVcpGznmhbWhwcVVgzTrDrDDLZlTLLZQrJ\nnhbnbfjcnfMfFNVtBq\nQHmPNZvfCLsSwJSm\npDhjpVDFcRBpFFjjMnRcVhpFCzbzsZbSSCtwtLMSLZLLtLbs\nFRjrnRchnfHPrrZlHl\nbjjMbdChgRDZthpQpRQnwRTprRwS\nmGzJsGsHzHGPvvvqvzGzSnglSJrrwQgnlQQSrlQl\nqvzHqHLHmHgPsNBdCdZtVBtVVMBFbh\nCctrCwrdpTwcpVrdpTpcrcnSJQttvQPHJQNQnQNjvvHQ\nzsqRlslRLqfgRmWsRgRzqzQnHjSBSQWJHPhHnSvHnJJJ\nGRgllbgfRgbzfRmwwcGdFMcTVPrFCF"
