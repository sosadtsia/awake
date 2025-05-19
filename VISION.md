# Awake: Project Vision

## Purpose
Awake exists to solve a common problem faced by macOS users: the need to temporarily prevent their Mac from entering sleep mode during important tasks. By providing a simple, intuitive command-line interface to macOS's built-in `caffeinate` functionality, Awake empowers users to maintain system wakefulness precisely when needed.

## Core Principles

1. **Simplicity**: Awake should remain easy to use with minimal cognitive overhead. The interface should be intuitive even for users with limited command-line experience.

2. **Reliability**: As a utility that users depend on for critical tasks (presentations, downloads, renders), Awake must be rock-solid reliable.

3. **Efficiency**: Awake should have minimal resource overhead and work seamlessly within the macOS ecosystem.

4. **Focused Functionality**: While we may expand features, each addition must directly serve the core purpose of controlling system sleep behavior.

## Current State

Awake currently provides:
- Basic sleep prevention with unlimited or time-limited duration
- Background operation mode for "set and forget" usage
- Various logging levels for different use cases
- Clean shutdown handling and signal management

## Strategic Roadmap

### Long-term Vision (1+ years)
- Support scheduled operation (e.g., "run every day between 2 PM and 4 PM")
- Add custom power management profiles (e.g., prevent only display sleep but allow system sleep)
- Develop a simple, native macOS GUI application as an alternative interface
- Implement event-based triggers (e.g., "stay awake while network is active")
- Explore cross-platform compatibility for Windows and Linux
- Build ecosystem integrations with popular productivity and automation tools

## Technology Choices

- **Go**: Chosen for its simplicity, performance, and cross-platform capabilities, long term support and backward compatibility.
- **Native macOS APIs**: Leverage system capabilities through the `caffeinate` command
- **Minimal Dependencies**: Maintain a small, focused dependency footprint

## Community and Contribution

We envision Awake as a community-driven project where users can easily contribute improvements. By maintaining clear documentation, well-structured code, and a welcoming contribution process, we aim to foster a collaborative development environment.

## Success Metrics

- **User Adoption**: Growth in downloads and active installations
- **Reliability**: Minimal bug reports and high stability
- **Contribution**: Active community participation and pull requests
- **Usefulness**: Positive user feedback and testimonials

---

This vision document is a living artifact and will evolve as the project matures and user needs change. The core purpose—helping users control their Mac's sleep behavior—remains constant, while implementation details may adapt to best fulfill that purpose.
