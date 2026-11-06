# Downloads / انتشار فایل‌های حجیم

مخزن فعلی شامل مستندات و کاتالوگ متنی است. فایل اجرایی و آرشیو کامل باینری هنوز در GitHub قرار نگرفته‌اند.

دلیل: مسیر اتصال GitHub موجود در این گفتگو، عملیات GitHub Contents برای فایل‌های متنی را در اختیار می‌گذارد و امکان آپلود مستقیم آرشیو ۱۱۷ تا ۱۲۶ مگابایتی یا فایل EXE به Release را ارائه نمی‌کند.

## روش انتشار از کامپیوتر صاحب مخزن

1. GitHub CLI را نصب کنید: https://cli.github.com/
2. در PowerShell اجرا کنید: `gh auth login`
3. فایل ZIP کامل و فایل EXE را روی کامپیوتر خود نگه دارید.
4. اسکریپت `tools/UPLOAD_RELEASE_WINDOWS.ps1` را از همین مخزن دریافت کنید.
5. اجرا کنید:

```powershell
.\UPLOAD_RELEASE_WINDOWS.ps1 -ZipPath "C:\path\JADDEH1405_v13_AllVehicles_Full.zip" -ExePath "C:\path\JADDEH1405_v13_AllVehicles.exe"
```

اسکریپت Release با برچسب `v1.3` ایجاد می‌کند و هر دو فایل را به‌عنوان Release asset آپلود می‌کند. برای انتشار بعدی می‌توانید `-Tag` را عوض کنید.

## لینک‌های مخزن

- [README فارسی](README_FA.md)
- [کاتالوگ ۴۷ وسیله نقلیه CSV](Models/Vehicle_Catalog_47.csv)
- [اسکریپت انتشار Release](tools/UPLOAD_RELEASE_WINDOWS.ps1)
