from urllib2 import urlopen
from bs4 import BeautifulSoup
import sys

class Core_Base:
    def __init__(self):
        self.word = ""
        self.part = ""
        self.mean =""
        self.trg = ""
        self.ind = ""
        self.p = ""
        self.ex = ""
        self.example =""
        self.mean_ex = {}
        self.subsense =""
        self.submean = ""
        self.subex = ""
        self.submean_ex = {}
        self.sub_len = 0

    #checks for every word using loop
    def words_p(self):
        self.format ="https://en.oxforddictionaries.com/definition/" + self.word
        self.search_for_meaning(self.format)
        
    def pass_words(self):
        self.data = {}
        self.word = sys.argv[1]
        self.words_p()
    
    #searches and stores the meaning of the particular word
    def search_for_meaning(self,url):
        html = urlopen(url)
        bsobj = BeautifulSoup(html.read(), 'lxml')
        self.section = bsobj.findAll("section", {"class": "gramb"})

        self.data[self.word] = [] 
                                                       
        #collects the meaning of the word from the web
        for i in self.section:
            self.span = i.find_all("span",{"class" : "pos"})
            for name in self.span:
                self.part = name.get_text()
            self.trg = i.find_all("div", {"class":"trg"})
            self.mean_ex[self.part] = []
            for j in self.trg:
                self.p = j.find_all("p")
                self.mean = ""
                for k in self.p:
                    self.ex =""
                    self.ind = k.find_all("span", {"class": "ind"})
                    for details in self.ind:
                        self.mean = (details.get_text())
                    self.ex = j.find("li", {"class" : "ex"})
                    self.example = str(self.ex)       
                    self.example = self.example[21:-12] 

                    if(len(self.example) == 0):
                        self.ex = j.find("div", {"class" : "ex"})
                        self.example = str(self.ex)       
                        self.example = self.example[22:-12]  
                    
                    self.mean_ex[self.part].append({                           
                        "meaning": self.mean,
                        "example": self.example
                    })
                    self.ex = ""
                    self.submean_ex[self.part] = []
                    self.subsense = j.find_all("li", {"class" : "subSense"})
                    for details in self.subsense:
                        self.submean = details.find("span", {"class" : "ind"})
                        self.submean = str(self.submean)
                        self.submean = self.submean[18:-7]
                        self.ex = details.find("li", {"class" : "ex"})
                        self.subex = str(self.ex)
                        self.subex = self.subex[21:-12]

                        if(len(self.subex) == 0):
                            self.ex = details.find("div", {"class" : "ex"})
                            self.subex = str(self.ex)
                            self.subex = self.subex[22:-12]

                        if(len(self.submean) !=0 and len(self.subex) != 0):
                            self.submean_ex[self.part].append({
                                "subMeaning": self.submean,
                                "subExample": self.subex
                            })
                    self.sub_len = (len(self.submean_ex[self.part]))    
                    if(self.sub_len != 0):
                        self.mean_ex[self.part].append(self.submean_ex[self.part])
                self.mean = ""
        self.data[self.word].append(self.mean_ex)                           
                
        self.data_len = len(self.data[self.word])
        if self.data_len < 1:
            del self.data[self.word]

    def display(self):

        mean = str(self.data[self.word])
        new = mean.replace("\\xe2",'')
        new = new.replace("\\x80",'')
        new = new.replace("\\x98",'')
        new = new.replace("\\x99",'')
        new = new.replace("u'", "'")
        new = new.replace("{'", "{*")
        new = new.replace("':", "*:")
        new = new.replace(" '", " *")
        new = new.replace(".'", ".*")
        new = new.replace('"', "*")

        print("word ", self.word )
        print(new)
        exit(0)

obj = Core_Base()
obj.pass_words()
obj.display()