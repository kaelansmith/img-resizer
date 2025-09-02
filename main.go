package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
)

const (
	maxLandscapeWidth  = 2500
	maxPortraitHeight  = 1500
	baseOutputDirName  = "Resized"
	supportedExtensions = ".jpg .jpeg .png .gif"
)

func main() {
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to get working directory:", err)
	}

	outputDir := filepath.Join(rootDir, baseOutputDirName)
	
	// Track skipped files per directory
	skippedCounts := make(map[string]int)
	err = filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		
		// Handle directories
		if d.IsDir() {
			if strings.HasPrefix(filepath.Base(path), baseOutputDirName) {
				// Skip already generated output folders
				return filepath.SkipDir
			}
			// Print directory entry with indentation based on depth
			relPath, _ := filepath.Rel(rootDir, path)
			depth := strings.Count(relPath, string(os.PathSeparator))
			indent := strings.Repeat("   ", depth)
			fmt.Printf("%sOpening sub-directory: %s\n", indent, filepath.Base(path))
			return nil
		}

		ext := strings.ToLower(filepath.Ext(d.Name()))
		if !strings.Contains(supportedExtensions, ext) {
			return nil
		}

		// Compute output path, preserving relative structure
		relPath, err := filepath.Rel(rootDir, path)
		if err != nil {
			return err
		}
		outPath := filepath.Join(outputDir, relPath)
		
		// Calculate indentation based on directory depth
		depth := strings.Count(relPath, string(os.PathSeparator))
		indent := strings.Repeat("   ", depth)
		
		// Check if output file already exists BEFORE processing
		if _, err := os.Stat(outPath); err == nil {
			// Get the directory path for grouping
			dirPath := filepath.Dir(relPath)
			if dirPath == "." {
				dirPath = "root"
			}
			skippedCounts[dirPath]++
			return nil
		}

		img, err := imaging.Open(path)
		if err != nil {
			fmt.Printf("%sSkipping (could not open): %s\n", indent, relPath)
			return nil
		}

		bounds := img.Bounds()
		width := bounds.Dx()
		height := bounds.Dy()

		// Determine if we need to resize
		var resized image.Image
		var needsResize bool
		if width > height && width > maxLandscapeWidth {
			resized = imaging.Resize(img, maxLandscapeWidth, 0, imaging.Lanczos)
			needsResize = true
		} else if height >= width && height > maxPortraitHeight {
			resized = imaging.Resize(img, 0, maxPortraitHeight, imaging.Lanczos)
			needsResize = true
		} else {
			resized = img // Use original image if no resize needed
			needsResize = false
		}
		
		if err := os.MkdirAll(filepath.Dir(outPath), os.ModePerm); err != nil {
			return err
		}

		if err := imaging.Save(resized, outPath); err != nil {
			return fmt.Errorf("failed to save image: %w", err)
		}

		if needsResize {
			fmt.Printf("%sResized: %s\n", indent, relPath)
		} else {
			fmt.Printf("%sCopied: %s\n", indent, relPath)
		}
		return nil
	})

	if err != nil {
		log.Fatal("Error walking directory:", err)
	}

	// Print summary of skipped files by directory
	if len(skippedCounts) > 0 {
		fmt.Println("\n📋 Skipped files summary:")
		for dirPath, count := range skippedCounts {
			if dirPath == "root" {
				fmt.Printf("   Root directory: %d files already exist\n", count)
			} else {
				fmt.Printf("   %s: %d files already exist\n", dirPath, count)
			}
		}
	}

	fmt.Println("\n✅ Done! Resized images are in:", outputDir)
	fmt.Println("\nPress Enter to close...")
	fmt.Scanln() // Wait for user to press Enter
}


