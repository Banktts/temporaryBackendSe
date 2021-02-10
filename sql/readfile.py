with open('/Users/thanabodeethamasaroch/go/src/Backend/sql/menu.csv') as f: #read file
    lines = f.readlines()
a=open("/Users/thanabodeethamasaroch/go/src/Backend/sql/mockData.sql","a") # write file
for i in lines:
    tmpDat=[]
    tmp=i.replace("\n","").split(",")
    for j in range(0,5):   #range of col change every time when change table
        if j==2 or j==4 or j==5: #number of col should be string
            tmpDat.append("\""+tmp[j]+"\"")
        else:
            tmpDat.append(tmp[j])

    a.write("("+",".join(tmpDat)+"),\n")
f.close()
a.close()
