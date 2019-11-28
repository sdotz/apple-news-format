# apple-news-format

A collection of structs and helpers for generating valid Apple News articles. Use this between your CMS and content model to assemble an article and serialize to the [Apple News Format (ANF)](https://developer.apple.com/documentation/apple_news/apple_news_format) JSON.

Publisher organizations have differing CMS systems and content formats, but you can use this to convert from JSON, XML or whatever structured formaty you like into Apple News format with a small amount of code.

It is common for CMSs to represent at least the body of the article as semantic HTML, so the library also includes helpers for converting from HTML to ANF autmatically. 

## Usage
```go
a := article.NewArticleWithDefaults()

a.Version = article.FORMAT_VERSION_1_7
a.Identifier = "newsjam://article/1c2ad8ac-9379-488a-9480-7f55b6da4a90" 
a.Language = "en-US"
a.DocumentStyle.BackgroundColor = "#fcfcfc"
a.Title := "My News Story"

bodyComponent := components.NewBody()
bodyComponent.SetText("This is the beginning of an amazing tale...")

// add a bunch of components to the Article
article.Components = append(article.Components, bodyComponent) 

// Article implements Stringer and will output the valid Apple News format JSON
fmt.Println(article.String())
```

## List of supported ANF components

### Text
|Component Type|Struct Support|HTML Conversion|
|---|---|---|
|Body|x|p|
|Title|x||
|Heading1,2,3,4,5,6|x|h1|
|ArticleTitle||
|Intro|x||
|Caption|x|figcaption|
|Author|x||
|Byline|x||
|Illustrator|x||
|Photographer|x||
|Quote|x||
|PullQuote|x||

### Images
|Component Type|Struct Support|HTML Conversion| 
|---|---|---|
|Image|x|img|
|Photo|x||
|Figure|x||
|Portrait|x||
|Logo|x||
|ArticleThumbnail|x||

### Galleries
|Component Type|Struct Support|
|---|---|
|Gallery|x|
|Mosaic|x|

### Audio
|Component Type|Struct Support|
|---|---|
|Audio|x|
|Music|x|
|Video|x|
|EmbedWebVideo|x|

### Location
|Component Type|Struct Support|
|---|---|
|Map|x|
|Place|x|

### Social
|Component Type|Struct Support|
|---|---|
|Instagram||
|FacebookPost||
|Tweet|x|

### Tables
|Component Type|Struct Support|
|---|---|
|DataTable||
|HTMLTable||

### Advertisements
|Component Type|Struct Support|
|---|---|
|BannerAdvertisement||
|MediumRectangleAdvertisement||
|ReplicaAdvertisement||

### Article Structure
|Component Type|Struct Support|
|---|---|
|Container|x|
|Section||
|Chapter||
|Aside||
|Header||
|Divider|x|
|ArticleLink||
|LinkButton||

### Augmented Reality
|Component Type|Struct Support|
|---|---|
|ARKit||
