# Image Resizer

A simple, fast Go-based image resizing utility that automatically resizes images while preserving their aspect ratios and directory structure. It's purposely opinionated for web devs building marketing sites, helping you go from overly-large high-res images to web-friendly sizes in no time.

## What It Does

This script processes all images in a directory (and its subdirectories) and creates a `Resized` folder containing:

- **Resized images** that exceed size limits (landscape > 2500px wide, portrait > 1500px tall)
- **Copied images** that are already within size limits
- **Preserved directory structure** - maintains the same folder hierarchy as the original

## Features

- **Automatic size detection** - Only resizes images that are too large
- **Aspect ratio preservation** - Maintains original proportions
- **Recursive processing** - Handles nested folders automatically
- **Incremental processing** - Skips files that have already been processed
- **Visual progress feedback** - Shows what's happening in real-time
- **High-quality resizing** - Uses Lanczos resampling for best results
- **Multiple format support** - Works with JPG, JPEG, PNG, and GIF files

## Size Limits

- **Landscape images**: Maximum width of 2500 pixels
- **Portrait images**: Maximum height of 1500 pixels
- **Square images**: Follows the larger dimension limit

## How to Use

### Step 1: Download the Right Executable

Choose the version for your system:

- **Windows**: [img-resizer-windows-amd64.exe](https://github.com/kaelansmith/img-resizer/releases/latest/download/img-resizer-windows-amd64.exe)
- **Mac Intel**: [img-resizer-darwin-amd64](https://github.com/kaelansmith/img-resizer/releases/latest/download/img-resizer-darwin-amd64)
- **Mac Apple Silicon**: [img-resizer-darwin-arm64](https://github.com/kaelansmith/img-resizer/releases/latest/download/img-resizer-darwin-arm64)
- **Linux**: [img-resizer-linux-amd64](https://github.com/kaelansmith/img-resizer/releases/latest/download/img-resizer-linux-amd64)

### Step 2: Copy to Your Image Folder

Move the downloaded executable to the folder containing your images.

### Step 3: Run It

**Windows:** Double-click the `.exe` file
**Mac/Linux:** Open Terminal/Command Prompt in your image folder and run:

```bash
./img-resizer-darwin-amd64  # or whatever version you downloaded
```

### Step 4: Wait for Completion

- A window will show the progress
- You'll see real-time updates as files are processed
- When done, press Enter to close

### Step 5: Find Your Results

Look for the `Resized` folder containing all your processed images.

## Example Output

```
Opening sub-directory: Photos
  Resized: Photos/vacation.jpg
  Copied: Photos/small-image.png
  Opening sub-directory: Photos/2023
    Resized: Photos/2023/large-photo.jpg
  Opening sub-directory: Photos/2024
    Copied: Photos/2024/small-photo.jpg

📋 Skipped files summary:
   Photos: 5 files already exist

✅ Done! Resized images are in: C:\path\to\Resized

Press Enter to close...
```

## File Structure

**Before:**

```
My Photos/
├── vacation.jpg (4000x3000)
├── small-image.png (800x600)
├── 2023/
│   └── large-photo.jpg (3000x2000)
└── 2024/
    └── small-photo.jpg (1200x800)
```

**After:**

```
My Photos/
├── Resized/
│   ├── vacation.jpg (2500x1875) ← Resized
│   ├── small-image.png (800x600) ← Copied
│   ├── 2023/
│   │   └── large-photo.jpg (2500x1667) ← Resized
│   └── 2024/
│       └── small-photo.jpg (1200x800) ← Copied
├── vacation.jpg (original)
├── small-image.png (original)
└── ... (all original files unchanged)
```

## Supported Formats

- **JPG/JPEG** - Most common photo format
- **PNG** - Good for graphics and screenshots
- **GIF** - Animated and static images

## Smart Features

### Incremental Processing

- **First run**: Processes all images
- **Subsequent runs**: Only processes new or changed images
- **Skipped files**: Shows summary of files that already exist

### Directory Handling

- **Skips "Resized" folders** - Won't process its own output
- **Visual hierarchy** - Shows folder structure with indentation
- **Error handling** - Gracefully handles unreadable files

### Performance

- **Fast processing** - Optimized for large image collections
- **Memory efficient** - Processes one image at a time
- **Progress feedback** - Real-time updates on what's happening

## Requirements

- **Windows, Mac, or Linux** - Download the appropriate executable for your system
- **No installation** - Just download and run
- **No dependencies** - Self-contained executable

## Tips

1. **Backup first** - Always keep your original images safe
2. **Test on a small folder** - Try it on a few images first
3. **Check the output** - Verify the results look good
4. **Run multiple times** - Safe to run repeatedly on the same folder
5. **Large collections** - Works great with thousands of images

## Troubleshooting

### "Could not open" errors

- File might be corrupted or in an unsupported format
- Check file permissions
- Try opening the file in another program

### No output folder created

- Check if you have write permissions in the directory
- Ensure the script completed successfully

### Slow performance

- Large images take longer to process
- Many files will take more time
- This is normal - the script is thorough

## Technical Details

- **Language**: Go (cross-platform)
- **Image library**: github.com/disintegration/imaging
- **Resampling**: Lanczos (high quality)
- **Output format**: Same as input format
- **Compression**: Maintains original quality settings

---

**Note**: This tool is designed to be simple and safe. It never modifies your original files and always creates a separate output folder.
