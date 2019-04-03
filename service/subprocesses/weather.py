import random
from urllib2 import urlopen
from bs4 import BeautifulSoup
import sys
import datetime
# from __future__ import print_function
pos=0
class Core_Base:
    def __init__(self):
        self.temperature = 0
        self.city = ""
        self.state = ""
        self.pos_of_temp=0
        self.pos2 = 0
        self.min_temp=0;self.max_temp=0; self.avg_temp = 0.0
        self.max_rain_probability=0.0; self.mix_rain_probability=0.0; self.avg_rain = 0.0

    def asking(self):
        self.city = sys.argv[1]
        self.state = sys.argv[2]
        self.country = "india"
        # url format design below
        self.format = "https://www.msn.com/en-in/weather/today/"+ self.city+self.state+self.country+"/we-city?q="+self.city+"-"+self.state+"&form=PRWLAS&iso=IN&el"
        self.extracting_info_from_net(self.format)


    # def is_float(self, input):
    #       try:
    #         # num = float(input)
    #       except ValueError:
    #         return False
    #       return True

    def is_int(self, input):
          try:
            num = int(input)
          except ValueError:
            return False
          return True

    def extracting_info_from_net(self, url):
        self.response = urlopen(url)
        self.source = self.response.read()
        self.source_text = BeautifulSoup(self.source)
        self.temp = self.source_text.text
        self.source_text = self.temp
        self.length_source_text = len(self.source_text)
        search_class = "Places"
        #print("search class = " + search_class)
        length_class = len(search_class)
        #print(self.source_text)
        c = 0
        for i in range(0, self.length_source_text - length_class):
            sub = self.source_text[i:i+length_class]
            if sub == search_class :
                # print('working')
                self.pos_of_temp = i + length_class

        #print("self.pos_of_temp : "+ str(self.pos_of_temp))
        # finding the first number after the statename
        for i in range(0,30):
            if self.is_int(self.source_text[i+self.pos_of_temp]):
                self.pos2 = i+self.pos_of_temp
                break


        if self.source_text[self.pos2 + 1].isnumeric() :
            #print("got it")
            self.temperature = int(self.source_text[self.pos2] + self.source_text[self.pos2 + 1])

        else :
            self.temperature = int(self.source_text[self.pos2])

    def further_info(self):
        # declaring and initializing

        self.feel = 0
        self.barometer = 0
        self.visibility = 0
        self.humidity = 0
        self.dew_point = 0
        l_feel = len("Feels") # where all l_<feature> represents the length of various parameters
        l_barometer = len("Barometer")
        l_visibility = len("Visibility")
        l_humidity = len("Humidity")
        l_dew_point = len("Dew Point")
        count = self.pos2 +2

        #working on these declared variables
        for i in range(count, count + 100):
            sub_1 = self.source_text[i:i + l_feel]
            sub_2 = self.source_text[i:i + l_barometer]
            sub_3 = self.source_text[i:i + l_visibility]
            sub_4 = self.source_text[i:i + l_humidity]
            sub_5 = self.source_text[i:i + l_dew_point]

            if sub_1 == "Feels":
                for j in range(0, 20):
                    if self.is_int(self.source_text[i + j]):
                        if self.source_text[i + j + 1].isnumeric() :
                            self.feel = int(self.source_text[i + j] + self.source_text[i + j + 1])

                        else :
                            self.feel = int(self.source_text[i + j])
                        break

            '''if sub_2 == "Barometer":
                for j in range(0, 20):
                    if self.is_int(self.source_text[i + j]):
                        if self.source_text[i + j + 1].isnumeric() :
                            self.barometer = int(self.source_text[i + j] + self.source_text[i + j + 1])

                        else :
                            self.barometer = int(self.source_text[i + j])
                        break'''
            # for present, Barometer part is skipped.
            if sub_3 == "Visibility":
                for j in range(0, 20):
                    if self.is_int(self.source_text[i + j]):
                        if self.source_text[i + j + 1].isnumeric() :
                            self.visibility = int(self.source_text[i + j] + self.source_text[i + j + 1])

                        else :
                            self.visibility = int(self.source_text[i + j])
                        break
            if sub_4 == "Humidity":
                for j in range(0, 20):
                    if self.is_int(self.source_text[i + j]):
                        if self.source_text[i + j + 1].isnumeric() :
                            self.humidity = int(self.source_text[i + j] + self.source_text[i + j + 1])

                        else :
                            self.humidity = int(self.source_text[i + j])
                        break
            if sub_5 == "Dew Point":
                for j in range(0, 20):
                    if self.is_int(self.source_text[i + j]):
                        if self.source_text[i + j + 1].isnumeric() :
                            self.dew_point = int(self.source_text[i + j] + self.source_text[i + j + 1])

                        else :
                            self.dew_point = int(self.source_text[i + j])
                        break



    def displaying(self):
        print("in json format")
        json = {
            "time": str(datetime.datetime.now()),
            "temperature": str(self.temperature),
            "city": self.city,
            "feels_like": str(self.feel),
            "visibility": str(self.visibility) + " km",
            "humidity": str(self.humidity) + " %",
            "dew_point": str(self.dew_point)
        }
        a = {}
        a["time"] = str(datetime.datetime.now())
        print(str(json).replace("'", '"'))
        print(a)
        exit(0)

obj = Core_Base()
obj.asking()
obj.further_info()
obj.displaying()