# JADDEH 1405

**Iranian open-road driving / exploration game project**

This repository contains the current development snapshot and documentation for JADDEH 1405.

## Project direction

- A road-trip and vehicle-interaction experience inspired by the *genre and gameplay loop* of open-road driving games.
- Iranian setting: roads, roadside stops, local vehicles, signs and everyday traffic.
- Vehicle catalogue and separate 3D asset pipeline.
- First-person on-foot exploration, vehicle entry/exit, repair and part interaction.
- Local save/load and Windows standalone build targets.

## Current snapshot

The previously prepared v1.3 package is a **prototype/development build**, not a finished photorealistic commercial game. Its vehicle catalogue contains 47 entries, but most vehicle meshes are procedural approximations. The Pride 131 asset has a more detailed separate-part model, but it is not a factory scan and should not be described as exact.

The Windows executable was cross-compiled, but it has not been verified by running it on a real Windows machine. Treat it as an unverified build until tested.

## Planned repository layout

- `Source/` — game source and build helpers
- `Models/` — vehicle catalogue and model exports
- `Pride131_Reference_Asset/` — Pride 131 model source and part metadata
- `Previews/` — screenshots and contact sheets
- `Builds/` — standalone Windows builds (when uploadable and tested)

## Requirements / notes

- The player-facing Windows build is intended to run without installing the game engine.
- Development and asset-generation tools are separate from the player build.
- Use only 3D assets, textures, sounds and reference data that are original or properly licensed.
- Do not redistribute third-party vehicle scans or model packs without permission.

## Persian

جاده ۱۴۰۵ یک پروژه شبیه‌ساز سفر جاده‌ای با فضای ایرانی است. بسته فعلی نسخه توسعه‌ای است و هنوز نسخه نهایی فوتورئالیستی محسوب نمی‌شود. مدل بیشتر خودروهای کاتالوگ بازسازی هندسی هستند؛ مدل پراید ۱۳۱ جزئیات بیشتری دارد اما اسکن کارخانه‌ای نیست. فایل اجرایی ویندوز نیز باید روی ویندوز واقعی آزمایش شود.

## License

No license is granted by this repository unless a license file is added. Third-party trademarks and vehicle designs remain the property of their respective owners.
