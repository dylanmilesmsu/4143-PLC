## P04
### Dylan Miles
### Description:
a program that concurrently downloads a set of images from given URLs and saves them to disk. By comparing the time taken to download images sequentially vs. concurrently, you will observe the benefits of concurrency for I/O-bound tasks.
### Files

|   #   | File            | Description                                        |
| :---: | --------------- | -------------------------------------------------- |
|   1   | [Main.go](./Main.go)         | Main driver of my project, contains all logic      |
|   2   | [go.mod](./go.mod)         | go module file      |
|   3   | [images.txt](./images.txt)         | example images file      |

### Realized performance gain


    Downloaded: Seq_image1.jpg
    Downloaded: Seq_image2.jpg
    Downloaded: Seq_image3.jpg
    Downloaded: Seq_image4.jpg
    Downloaded: Seq_image5.jpg
    Sequential download took: 1.042331242s
    Downloaded: Async_image2.jpg
    Downloaded: Async_image1.jpg
    Downloaded: Async_image5.jpg
    Downloaded: Async_image4.jpg
    Downloaded: Async_image3.jpg
    Concurrent download took: 50.934761ms

### Instructions

- Example run command:
    - `go run "./Main.go`
