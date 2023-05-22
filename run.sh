ffmpeg -i input.mp4 -ss 00:00:00 -to 00:10:00 -c copy dist/output1.mp4
ffmpeg -i input.mp4 -ss 00:10:00 -to 00:20:00 -c copy dist/output2.mp4
ffmpeg -i input.mp4 -ss 00:20:00 -to 00:30:00 -c copy dist/output3.mp4