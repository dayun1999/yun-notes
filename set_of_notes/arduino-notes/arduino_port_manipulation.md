# Arduino port manipulation (Arduino的端口操作)

## 推荐文章

- [Arduino Port-Manipulation von Prof. Jürgen Plate](http://www.netzmafia.de/skripten/hardware/Arduino/Programmierung/portmanipulation.html)

## 1.bitwise 位操作

```c
#define bitRead(value, bit) (((value) >> (bit)) & 0x01)

#define bitSet(value, bit) ((value) |= (1UL << (bit)))

#define bitClear(value, bit) ((value) &= ~(1UL << (bit)))

#define bitWrite(value, bit, bitvalue) ((bitvalue) ? bitSet(value, bit) : bitClear(value, bit))
```
