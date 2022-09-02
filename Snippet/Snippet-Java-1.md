## 删除Excels Sheet页
---
```java
for (Map.Entry<String, Object> entry : sheetMap.entrySet()) {
    String sheetName = entry.getKey();
    if (workbook.getSheet(sheetName).getPhysicalNumberOfRows() <= 2){
        workbook.removeSheetAt(workbook.getSheetIndex(sheetName));
    }
}
```

## List
---
```java
// 删除null
list.removeIf(Objects::isNull);

// 截取子list
list.subList(3, 5)
```

## String
---
```java
// 切割String
StringUtils.delimitedListToStringArray(name, "=")
```

## HashMap
---
```java
// 初始化
Map<String, String> myMap = new HashMap<String, String>();  
myMap.put("a", "b");  
myMap.put("c", "d");

// 初始化java8
HashMap<String, String > myMap  = new HashMap<String, String>(){{  
      put("a","b");  
      put("b","b");       
}};

// 初始化java9
Map.of("Hello", 1, "World", 2);//不可变集合
```